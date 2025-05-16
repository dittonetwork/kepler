#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if script is run with sudo
if [ "$EUID" -ne 0 ]; then
    echo -e "${RED}[✗] This script must be run with sudo privileges${NC}"
    echo "Usage: sudo $0"
    exit 1
fi

# Function to print status messages
print_status() {
    echo -e "${GREEN}[✓] $1${NC}"
}

print_error() {
    echo -e "${RED}[✗] $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}[!] $1${NC}"
}

# Get the original user's home directory
if [ -n "$SUDO_USER" ]; then
    USER_HOME=$(getent passwd "$SUDO_USER" | cut -d: -f6)
else
    USER_HOME="$HOME"
fi

# Check if kepler is already initialized
if [ -d "$USER_HOME/.kepler" ]; then
    print_warning "Kepler is already initialized in $USER_HOME/.kepler"
    read -p "Do you want to reinitialize? This will remove the existing configuration. (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_status "Setup cancelled"
        exit 0
    fi

    # Stop and disable the service if it exists
    if systemctl is-active --quiet kepler.service; then
        print_status "Stopping Kepler service..."
        systemctl stop kepler.service
        systemctl disable kepler.service
    fi

    print_status "Removing existing Kepler configuration..."
    sudo -u "$SUDO_USER" rm -rf "$USER_HOME/.kepler"
fi

print_status "Starting Kepler validator setup..."

# Download and install keplerd
print_status "Downloading keplerd..."
wget https://github.com/dittonetwork/kepler/releases/download/v0.1.0/v0.1.0_linux_amd64.tar.gz || { print_error "Failed to download keplerd"; exit 1; }

print_status "Extracting keplerd..."
tar -xzf v0.1.0_linux_amd64.tar.gz -C . || { print_error "Failed to extract archive"; exit 1; }

print_status "Moving keplerd..."
mv tmp/4274431292/keplerd /usr/local/bin || { print_error "Failed to move keplerd"; exit 1; }
rm -rf ./tmp
rm v0.1.0_linux_amd64.tar.gz
chmod +x /usr/local/bin/keplerd

# Initialize keplerd
print_status "Initializing keplerd..."
sudo -u "$SUDO_USER" keplerd init "kepler" --chain-id testnet || { print_error "Failed to initialize keplerd"; exit 1; }

# Configure keplerd
print_status "Configuring keplerd..."
sudo -u "$SUDO_USER" sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.000025ditto"/' "$USER_HOME/.kepler/config/app.toml"
sudo -u "$SUDO_USER" sed -i 's/seeds = ""/seeds = "1fbe5abb67bcfaafb753da7f3e5a8ce9b04111c4@3.76.127.157:26656"/' "$USER_HOME/.kepler/config/config.toml"

# Download genesis file
print_status "Downloading genesis file..."
sudo -u "$SUDO_USER" curl -s https://static.dittonetwork.io/genesis.json > "$USER_HOME/.kepler/config/genesis.json" || { print_error "Failed to download genesis file"; exit 1; }

# Generate validator key
print_status "Setting up validator key..."
read -sp "Enter your 64-byte hex private key (128 hex characters) (input will be hidden): " HEX_PRIV
echo # Add a newline after the hidden input

if [ -z "$HEX_PRIV" ]; then
    print_error "Private key is required"
    exit 1
fi

if [ ${#HEX_PRIV} -ne 128 ]; then
    print_error "Private key must be 64 bytes (128 hex characters)"
    exit 1
fi

# Convert hex to raw bytes (requires xxd)
RAW_PRIV=$(echo "$HEX_PRIV" | xxd -r -p)

# Base64 encode the full private key
BASE64_PRIV=$(echo -n "$RAW_PRIV" | base64 | tr -d '\n')

# Extract public key (last 32 bytes)
BASE64_PUB=$(echo -n "$RAW_PRIV" | tail -c 32 | base64 | tr -d '\n')

# Calculate address = sha256(pubkey)[:20], then hex uppercase
ADDRESS=$(echo -n "$RAW_PRIV" | tail -c 32 | sha256sum | cut -c1-40 | tr 'a-f' 'A-F' | tr -d '\n')

# Create priv_validator_key.json
print_status "Creating validator key file..."
cat > "$USER_HOME/.kepler/config/priv_validator_key.json" <<EOF
{
  "address": "${ADDRESS}",
  "pub_key": {
    "type": "tendermint/PubKeyEd25519",
    "value": "${BASE64_PUB}"
  },
  "priv_key": {
    "type": "tendermint/PrivKeyEd25519",
    "value": "${BASE64_PRIV}"
  }
}
EOF

# Fix permissions for the validator key file
chown "$SUDO_USER:$SUDO_USER" "$USER_HOME/.kepler/config/priv_validator_key.json"
chmod 600 "$USER_HOME/.kepler/config/priv_validator_key.json"

# Create systemd service
print_status "Creating systemd service..."
cat > /etc/systemd/system/kepler.service << EOS
[Unit]
Description=Kepler Network
After=network.target

[Service]
User=$SUDO_USER
ExecStart=/usr/local/bin/keplerd start
Restart=always
StartLimitBurst=999
RestartSec=10

[Install]
WantedBy=multi-user.target
EOS

# Reload and start service
print_status "Starting Kepler service..."
systemctl daemon-reload
systemctl start kepler.service
systemctl enable kepler.service

# Verify service status
print_status "Verifying service status..."
if systemctl is-active --quiet kepler.service; then
    print_status "Kepler service is running successfully"
else
    print_error "Kepler service failed to start"
    exit 1
fi

print_status "Setup completed successfully!"
print_status "You can check the service status with: systemctl status kepler.service"
print_status "You can check the logs with: journalctl -u kepler.service -f"
