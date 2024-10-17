# Makefile for Kepler Testnet

# Variables
HOME_DIR := $$(pwd)
TESTNET_DIR := $(HOME_DIR)/testnet
NODE1_DIR := $(TESTNET_DIR)/node1
NODE2_DIR := $(TESTNET_DIR)/node2
NODE1_HOME := $(NODE1_DIR)/node1home
NODE2_HOME := $(NODE2_DIR)/node2home

# Targets
.PHONY: init clean setup build init-nodes add-genesis-accounts gentx collect-gentxs start-node1 start-node2 create-validator

init: clean setup build init-nodes add-genesis-accounts gentx collect-gentxs

clean:
	@rm -rf testnet

setup:
	@mkdir -p $(NODE1_HOME) $(NODE2_HOME)

build:
	@echo "Building Kepler chain..."
	@ignite chain build

init-nodes:
	@echo "Initializing nodes..."
	@keplerd init node1 --home $(NODE1_HOME)
	@keplerd init node2 --home $(NODE2_HOME)

add-genesis-accounts:
	@echo "Adding genesis accounts..."
	@yes | keplerd keys add node1-account --home $(NODE1_HOME)
	@yes | keplerd keys add node2-account --home $(NODE2_HOME)
	@keplerd genesis add-genesis-account $$(keplerd keys show node1-account -a --home $(NODE1_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)
	@keplerd genesis add-genesis-account $$(keplerd keys show node2-account -a --home $(NODE2_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)

gentx:
	@echo "Generating genesis transaction..."
	@keplerd genesis gentx node1-account 10000000000stake --fees 100000stake --home $(NODE1_HOME)

collect-gentxs:
	@echo "Collecting genesis transactions..."
	@keplerd genesis collect-gentxs --home $(NODE1_HOME) --trace

start-node1:
	@echo "Starting node1..."
	@awk '{gsub(/enable = false/, "enable = true"); gsub(/swagger = false/, "swagger = true"); if ($$0 ~ /^minimum-gas-prices =/) {$$0 = "minimum-gas-prices = \"0.1stake\""} print}' $(NODE1_HOME)/config/app.toml > $(NODE1_HOME)/config/app.toml.tmp && mv $(NODE1_HOME)/config/app.toml.tmp $(NODE1_HOME)/config/app.toml
	@keplerd start --home $(NODE1_HOME)

start-node2: copy-config
	@echo "Starting node2..."
	@awk '{gsub(/tcp:\/\/0.0.0.0:1317/, "tcp://0.0.0.0:1318"); gsub(/0.0.0.0:9090/, "0.0.0.0:9092"); gsub(/0.0.0.0:9091/, "0.0.0.0:9093"); print}' $(NODE2_HOME)/config/app.toml > $(NODE2_HOME)/config/app.toml.tmp && mv $(NODE2_HOME)/config/app.toml.tmp $(NODE2_HOME)/config/app.toml
	@awk '{gsub(/tcp:\/\/127.0.0.1:26658/, "tcp://127.0.0.1:26659"); gsub(/tcp:\/\/127.0.0.1:26657/, "tcp://127.0.0.1:26654"); gsub(/localhost:6060/, "localhost:6062"); gsub(/0.0.0.0:26656/, "0.0.0.0:26653"); print}' $(NODE2_HOME)/config/config.toml > $(NODE2_HOME)/config/config.toml.tmp && mv $(NODE2_HOME)/config/config.toml.tmp $(NODE2_HOME)/config/config.toml
	@awk '{gsub(/tcp:\/\/localhost:26657/, "tcp://localhost:26654"); print}' $(NODE2_HOME)/config/client.toml > $(NODE2_HOME)/config/client.toml.tmp && mv $(NODE2_HOME)/config/client.toml.tmp $(NODE2_HOME)/config/client.toml
	@NODE1_ID=$$(keplerd comet show-node-id --home $(NODE1_HOME)) && \
		awk -v node1_id=$$NODE1_ID '{gsub(/^seeds = ""/, "seeds = \"" node1_id "@127.0.0.1:26656\""); gsub(/^persistent_peers = ""/, "persistent_peers = \"" node1_id "@127.0.0.1:26656\""); print}' $(NODE2_HOME)/config/config.toml > $(NODE2_HOME)/config/config.toml.tmp && mv $(NODE2_HOME)/config/config.toml.tmp $(NODE2_HOME)/config/config.toml
	@keplerd start --home $(NODE2_HOME)

copy-config:
	@echo "Copying configuration from node1 to node2..."
	@cp $(NODE1_HOME)/config/{genesis.json,config.toml,app.toml,client.toml} $(NODE2_HOME)/config/

create-validator:
	@echo "Creating validator for node2..."
	@sh ./scripts/create_validator.sh $(NODE2_HOME) $(NODE2_DIR)

check-validators:
	@echo "Checking validators..."
	@keplerd q comet-validator-set --home $(NODE1_HOME)