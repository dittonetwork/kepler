#!/bin/bash

set -e

# 📍 Save current directory
SCRIPT_DIR="$(pwd)"

# 1. Navigate to Kepler config directory
CONFIG_DIR="$HOME/.kepler/config"
if [ ! -d "$CONFIG_DIR" ]; then
  echo "❌ Directory $CONFIG_DIR does not exist. Exiting."
  exit 1
fi

echo "📁 Navigating to $CONFIG_DIR"
cd "$CONFIG_DIR"

# 2. Edit app.toml
echo "⚙️ Updating app.toml..."
sed -i.bak '/\[telemetry\]/,/^\[/ s/enabled *= *.*/enabled = true/; s/prometheus-retention-time *= *.*/prometheus-retention-time = 180/' app.toml || \
echo -e "\n[telemetry]\nenabled = true\nprometheus-retention-time = 180" >> app.toml

# 3. Edit config.toml
echo "⚙️ Updating config.toml..."
sed -i.bak 's/^prometheus *= *.*/prometheus = true/; s/^prometheus_listen_addr *= *.*/prometheus_listen_addr = "0.0.0.0:26660"/' config.toml || \
echo -e "\nprometheus = true\nprometheus_listen_addr = \"0.0.0.0:26660\"" >> config.toml

# 4. Add to /etc/hosts if not already present
HOST_ENTRY="172.17.0.1 host.docker.internal"
if ! grep -qF "$HOST_ENTRY" /etc/hosts; then
  echo "🛠️ Adding host entry to /etc/hosts (requires sudo)"
  echo "$HOST_ENTRY" | sudo tee -a /etc/hosts > /dev/null
else
  echo "✅ Host entry already exists in /etc/hosts"
fi

# 🔙 Return to original script directory (kepler/monitor)
cd "$SCRIPT_DIR"

# 5. Start the containers
echo "🚀 Starting Docker containers..."
docker compose up -d

# 6. Check that everything is running correctly
echo "📊 Checking container status..."
docker compose ps
