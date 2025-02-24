BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
APPNAME := kepler
COVER_FILE := coverage.txt
COVER_HTML_FILE := coverage.html

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

# Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=$(APPNAME) \
	-X github.com/cosmos/cosmos-sdk/version.AppName=$(APPNAME)d \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

##############
###  Test  ###
##############

test-unit:
	@echo Running unit tests...
	@go test -mod=readonly -v -timeout 30m ./...

test-race:
	@echo Running unit tests with race condition reporting...
	@go test -mod=readonly -v -race -timeout 30m ./...

test-cover:
	@echo Running unit tests and creating coverage report...
	@go test -mod=readonly -v -timeout 30m -coverprofile=$(COVER_FILE) -covermode=atomic ./...
	@go tool cover -html=$(COVER_FILE) -o $(COVER_HTML_FILE)

bench:
	@echo Running unit tests with benchmarking...
	@go test -mod=readonly -v -timeout 30m -bench=. ./...

test: test-unit

.PHONY: test test-unit test-race test-cover bench

#################
###  Install  ###
#################

all: install

install:
	@echo "--> ensure dependencies have not been modified"
	@go mod verify
	@echo "--> installing $(APPNAME)d"
	@go install $(BUILD_FLAGS) -mod=readonly ./cmd/$(APPNAME)d

.PHONY: all install

##################
###  Protobuf  ###
##################

# Use this proto-image if you do not want to use Ignite for generating proto files
protoVer=0.15.1
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-gen:
	@echo "Generating protobuf files..."
	@ignite generate proto-go --yes

.PHONY: proto-gen

.PHONY: proto-lint
proto-lint:
	buf lint

.PHONY: proto-format
proto-format:
	buf format -w


#################
###  Linting  ###
#################

golangci_lint_cmd=golangci-lint
golangci_version=v1.64.2

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --timeout 15m

lint-fix:
	@echo "--> Running linter and fixing issues"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --fix --timeout 15m

.PHONY: lint lint-fix

###################
### Development ###
###################

govet:
	@echo Running go vet...
	@go vet ./...

govulncheck:
	@echo Running govulncheck...
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...

.PHONY: govet govulncheck

###############################
### Docker & Ignite Build  ####
###############################

AWS_REGION := ap-southeast-1
ECR_REPO := 847647377987.dkr.ecr.$(AWS_REGION).amazonaws.com/kepler:latest

ecr-login:
	@echo "--> Logging in to Amazon ECR"
	@aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin 847647377987.dkr.ecr.$(AWS_REGION).amazonaws.com

# Building binary file via Ignite Build for linux/amd64
ignite-build:
	@echo "--> Building binary with Ignite"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ignite chain build

# Creatiton of Docker-image. Dockerfile use.
docker-build:
	@echo "--> Building Docker image"
	@docker buildx build --platform linux/amd64 -t $(ECR_REPO) .

# Push into Amazon ECR
docker-push:
	@echo "--> Pushing Docker image to Amazon ECR"
	@docker push $(ECR_REPO)

build-all: ecr-login ignite-build docker-build docker-push

.PHONY: ecr-login ignite-build docker-build docker-push build-all