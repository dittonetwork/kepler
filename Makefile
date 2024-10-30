# Makefile for Kepler localnet

# Variables
HOME_DIR := $$(pwd)
LOCALNET_DIR := $(HOME_DIR)/localnet

NODE1_DIR := $(LOCALNET_DIR)/node1
NODE2_DIR := $(LOCALNET_DIR)/node2
NODE3_DIR := $(LOCALNET_DIR)/node3
NODE4_DIR := $(LOCALNET_DIR)/node4
NODE5_DIR := $(LOCALNET_DIR)/node5

NODE1_HOME := $(NODE1_DIR)/node1home
NODE2_HOME := $(NODE2_DIR)/node2home
NODE3_HOME := $(NODE3_DIR)/node3home
NODE4_HOME := $(NODE4_DIR)/node4home
NODE5_HOME := $(NODE5_DIR)/node5home

# Targets
.PHONY: init clean setup build start-node1 start-node2 create-validator up down localnet create-validators-localnet check-validators

init: clean setup build init-nodes add-genesis-accounts gentx collect-gentxs

localnet: down clean setup build build-docker init-nodes add-genesis-accounts gentx collect-gentxs update-configs up
localnet-postgres: down clean setup build build-docker init-nodes setup-postgres-indexer add-genesis-accounts gentx collect-gentxs update-configs up

clean:
	@rm -rf localnet

setup:
	@mkdir -p $(NODE1_HOME) $(NODE2_HOME) $(NODE3_HOME) $(NODE4_HOME) $(NODE5_HOME)

build:
	@echo "Building Kepler chain..."
	@ignite chain build

build-docker:
	@docker build -t kepler:local -f build/Dockerfile .

init-nodes:
	@echo "Initializing nodes..."

	@keplerd init node1 --home $(NODE1_HOME)
	@keplerd init node2 --home $(NODE2_HOME)
	@keplerd init node3 --home $(NODE3_HOME)
	@keplerd init node4 --home $(NODE4_HOME)
	@keplerd init node5 --home $(NODE5_HOME)

setup-postgres-indexer:
	@echo "Setting up postgres indexer..."

	@keplerd config set --home $(NODE1_HOME) config tx_index.indexer psql -s
	@keplerd config set --home $(NODE1_HOME) config tx_index.psql-conn postgresql://kepler:kepler@192.168.10.7:5432/kepler?sslmode=disable -s

add-genesis-accounts:
	@echo "Adding genesis accounts..."

	@yes "again glare leg choice input april tone brush goose seek forum dinosaur link speed digital question ticket caught strike quote release crowd fork deposit"     | keplerd keys add node1-account --recover --keyring-backend=test --home $(NODE1_HOME)
	@yes "heavy notice slice boil dune monitor pet slight denial sea train notable section boring fancy evolve kangaroo lazy pride cluster gaze any hotel chimney"       | keplerd keys add node2-account --recover --keyring-backend=test --home $(NODE2_HOME)
	@yes "crime salute atom that obey gaze among arrow decide bicycle robust glow diesel possible hill system three ethics share violin recipe oven dice height"         | keplerd keys add node3-account --recover --keyring-backend=test --home $(NODE3_HOME)
	@yes "follow appear body worry number laundry crucial fiction swing sure image poem alter hole motor chef traffic enact near arrange imitate shoulder saddle hybrid" | keplerd keys add node4-account --recover --keyring-backend=test --home $(NODE4_HOME)
	@yes "cancel welcome evolve effort luggage long foam teach cage betray library damp pink artist dilemma foster mean blast this moon charge lobster swarm ordinary"   | keplerd keys add node5-account --recover --keyring-backend=test --home $(NODE5_HOME)

	@keplerd genesis add-genesis-account $$(keplerd keys show node1-account -a --keyring-backend=test --home $(NODE1_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)
	@keplerd genesis add-genesis-account $$(keplerd keys show node2-account -a --keyring-backend=test --home $(NODE2_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)
	@keplerd genesis add-genesis-account $$(keplerd keys show node3-account -a --keyring-backend=test --home $(NODE3_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)
	@keplerd genesis add-genesis-account $$(keplerd keys show node4-account -a --keyring-backend=test --home $(NODE4_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)
	@keplerd genesis add-genesis-account $$(keplerd keys show node5-account -a --keyring-backend=test --home $(NODE5_HOME)) 100000000000000000000000000stake --home $(NODE1_HOME)

gentx:
	@echo "Generating genesis transaction..."
	@keplerd genesis gentx node1-account 10000000000stake --fees 100000stake --keyring-backend=test --home $(NODE1_HOME)

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

update-configs:
	@awk '{gsub(/enable = false/, "enable = true"); gsub(/swagger = false/, "swagger = true"); if ($$0 ~ /^minimum-gas-prices =/) {$$0 = "minimum-gas-prices = \"0.1stake\""} gsub(/127\.0\.0\.1/, "0.0.0.0"); gsub(/localhost/, "0.0.0.0"); print}' $(NODE1_HOME)/config/app.toml > $(NODE1_HOME)/config/app.toml.tmp && mv $(NODE1_HOME)/config/app.toml.tmp $(NODE1_HOME)/config/app.toml
	@awk '{gsub(/127\.0\.0\.1/, "0.0.0.0"); print}' $(NODE1_HOME)/config/config.toml > $(NODE1_HOME)/config/config.toml.tmp && mv $(NODE1_HOME)/config/config.toml.tmp $(NODE1_HOME)/config/config.toml

	@cp $(NODE1_HOME)/config/{genesis.json,config.toml,app.toml,client.toml} $(NODE2_HOME)/config/
	@cp $(NODE1_HOME)/config/{genesis.json,config.toml,app.toml,client.toml} $(NODE3_HOME)/config/
	@cp $(NODE1_HOME)/config/{genesis.json,config.toml,app.toml,client.toml} $(NODE4_HOME)/config/
	@cp $(NODE1_HOME)/config/{genesis.json,config.toml,app.toml,client.toml} $(NODE5_HOME)/config/

	@NODE1_ID=$$(keplerd comet show-node-id --home $(NODE1_HOME)) && \
		awk -v node1_id=$$NODE1_ID '{gsub(/^seeds = ""/, "seeds = \"" node1_id "@192.168.10.2:26656\""); gsub(/^persistent_peers = ""/, "persistent_peers = \"" node1_id "@192.168.10.2:26656\""); print}' $(NODE2_HOME)/config/config.toml > $(NODE2_HOME)/config/config.toml.tmp && mv $(NODE2_HOME)/config/config.toml.tmp $(NODE2_HOME)/config/config.toml && \
		awk -v node1_id=$$NODE1_ID '{gsub(/^seeds = ""/, "seeds = \"" node1_id "@192.168.10.2:26656\""); gsub(/^persistent_peers = ""/, "persistent_peers = \"" node1_id "@192.168.10.2:26656\""); print}' $(NODE3_HOME)/config/config.toml > $(NODE3_HOME)/config/config.toml.tmp && mv $(NODE3_HOME)/config/config.toml.tmp $(NODE3_HOME)/config/config.toml && \
		awk -v node1_id=$$NODE1_ID '{gsub(/^seeds = ""/, "seeds = \"" node1_id "@192.168.10.2:26656\""); gsub(/^persistent_peers = ""/, "persistent_peers = \"" node1_id "@192.168.10.2:26656\""); print}' $(NODE4_HOME)/config/config.toml > $(NODE4_HOME)/config/config.toml.tmp && mv $(NODE4_HOME)/config/config.toml.tmp $(NODE4_HOME)/config/config.toml && \
		awk -v node1_id=$$NODE1_ID '{gsub(/^seeds = ""/, "seeds = \"" node1_id "@192.168.10.2:26656\""); gsub(/^persistent_peers = ""/, "persistent_peers = \"" node1_id "@192.168.10.2:26656\""); print}' $(NODE5_HOME)/config/config.toml > $(NODE5_HOME)/config/config.toml.tmp && mv $(NODE5_HOME)/config/config.toml.tmp $(NODE5_HOME)/config/config.toml

up:
	@docker compose up -d

down:
	@docker compose down

create-validator:
	@echo "Creating validator for node2..."
	@sh ./scripts/create_validator.sh $(NODE2_HOME) $(NODE2_DIR) node2 node2-account

create-validators-localnet:
	@echo "Creating validators for localnet..."
	@sh ./scripts/create_validator.sh $(NODE2_HOME) $(NODE2_DIR) node2 node2-account
	@sh ./scripts/create_validator.sh $(NODE3_HOME) $(NODE3_DIR) node3 node3-account
	@sh ./scripts/create_validator.sh $(NODE4_HOME) $(NODE4_DIR) node4 node4-account
	@sh ./scripts/create_validator.sh $(NODE5_HOME) $(NODE5_DIR) node5 node5-account

check-validators:
	@echo "Checking validators..."
	@keplerd q comet-validator-set --home $(NODE1_HOME)
