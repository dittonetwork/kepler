#!/bin/bash

# Exit on error
set -e

# Check for protoc
if ! command -v protoc &> /dev/null; then
    echo "protoc is not installed"
    exit 1
fi

# Base directories
PROTO_PATH="./proto"
THIRD_PARTY_PROTO="./third_party/proto"
PROTO_FILES="${PROTO_PATH}/kepler/kepler/*.proto"

# Create necessary directories
mkdir -p "${THIRD_PARTY_PROTO}"
mkdir -p ./api

# Function to download and extract proto files
download_protos() {
    local REPO=$1
    local VERSION=${2:-main}
    local TARGET_DIR="${THIRD_PARTY_PROTO}/$3"
    local STRIP=${4:-1}

    echo "Downloading protos from ${REPO} (${VERSION})..."
    mkdir -p "${TARGET_DIR}"

    curl -L "https://github.com/${REPO}/archive/${VERSION}.tar.gz" | \
    tar -xz --strip-components=${STRIP} -C "${TARGET_DIR}"
}

echo "Downloading dependencies..."

# Download all required proto files with specific versions
download_protos "cosmos/cosmos-sdk" "v0.47.0" "cosmos-sdk"
download_protos "cosmos/cosmos-proto" "v1.0.0-beta.2" "cosmos-proto/proto" 2
download_protos "cosmos/gogoproto" "v1.4.1" "gogoproto"
#download_protos "cosmos/ics23" "v0.9.0" "ics23"
download_protos "cosmos/ibc-go" "v7.0.0" "ibc"
download_protos "googleapis/googleapis" "master" "googleapis"

# Organize proto files
#mkdir -p "${THIRD_PARTY_PROTO}/amino"
#mkdir -p "${THIRD_PARTY_PROTO}/cosmos_proto"
#mkdir -p "${THIRD_PARTY_PROTO}/cosmos"
#mkdir -p "${THIRD_PARTY_PROTO}/google"

# Copy required protos to their correct locations
#cp -r "${THIRD_PARTY_PROTO}/cosmos-sdk/proto/amino/"* "${THIRD_PARTY_PROTO}/amino/"
#cp -r "${THIRD_PARTY_PROTO}/cosmos-proto/proto/cosmos_proto/"* "${THIRD_PARTY_PROTO}/cosmos_proto/"
#cp -r "${THIRD_PARTY_PROTO}/cosmos-sdk/proto/cosmos/"* "${THIRD_PARTY_PROTO}/cosmos/"
#cp -r "${THIRD_PARTY_PROTO}/googleapis/google/"* "${THIRD_PARTY_PROTO}/google/"

# Create include paths string
INCLUDES="-I${PROTO_PATH}"
INCLUDES="${INCLUDES} -I${THIRD_PARTY_PROTO}"
INCLUDES="${INCLUDES} -I${THIRD_PARTY_PROTO}/cosmos-sdk/proto"
INCLUDES="${INCLUDES} -I${THIRD_PARTY_PROTO}/cosmos-sdk/third_party/proto"
INCLUDES="${INCLUDES} -I${THIRD_PARTY_PROTO}/gogoproto"
INCLUDES="${INCLUDES} -I${THIRD_PARTY_PROTO}/googleapis"

echo "Starting proto generation..."

# TypeScript generation
#echo "Generating TypeScript..."
#protoc ${INCLUDES} \
#    --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
#    --ts_proto_out=. \
#    --ts_proto_opt=logtostderr=true \
#    --ts_proto_opt=allow_merge=true \
#    --ts_proto_opt=json_names_for_fields=false \
#    --ts_proto_opt=snakeToCamel=true \
#    --ts_proto_opt=esModuleInterop=true \
#    ${PROTO_FILES}
#
## Swagger/OpenAPI generation
#echo "Generating Swagger..."
#protoc ${INCLUDES} \
#    --openapiv2_out=. \
#    --openapiv2_opt=logtostderr=true \
#    --openapiv2_opt=openapi_naming_strategy=fqn \
#    --openapiv2_opt=json_names_for_fields=false \
#    --openapiv2_opt=generate_unbound_methods=true \
#    ${PROTO_FILES}

# Gogo generation
echo "Generating Gogo..."
protoc ${INCLUDES} \
    --gocosmos_out=plugins=grpc:../ \
    --gocosmos_opt=Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types \
    --gocosmos_opt=Mcosmos/orm/v1/orm.proto=cosmossdk.io/orm \
    --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:../ \
    ${PROTO_FILES}

# STA generation
#echo "Generating STA..."
#protoc ${INCLUDES} \
#    --openapiv2_out=. \
#    --openapiv2_opt=logtostderr=true \
#    --openapiv2_opt=openapi_naming_strategy=simple \
#    --openapiv2_opt=ignore_comments=true \
#    --openapiv2_opt=simple_operation_ids=false \
#    --openapiv2_opt=json_names_for_fields=false \
#    ${PROTO_FILES}

# Pulsar generation
echo "Generating Pulsar..."
protoc ${INCLUDES} \
    --go-pulsar_out=paths=source_relative:./api \
    --go-grpc_out=paths=source_relative:./api \
    ${PROTO_FILES}

echo "Proto generation completed!"

# Optional: Remove temporary files
# rm -rf "${THIRD_PARTY_PROTO}"