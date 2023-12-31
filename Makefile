GEN_DIR=proto-gen
proto:
	 protoc --go_out=${GEN_DIR} --go_opt=paths=import \
        -I=./idl/proto/api_v1/ \
		--go-grpc_out=${GEN_DIR} --go-grpc_opt=paths=source_relative \
		idl/proto/api_v1/*.proto

proto-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

build:
	go build -ldflags "-X 'main.GoVersion=$(shell go version)' -X 'main.GitHash=$(shell git show -s --format=%H)' -X 'main.BuildTime=$(shell git show -s --format=%cd)'" -o collector ./cmd/collector/collectorserver.go
