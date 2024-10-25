# Install protoc plugins
go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@latest
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest

npm install ts-proto