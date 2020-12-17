#!/bin/sh

go mod tidy && go install \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

for f in $(find . -name "*.proto" -not -path "*internal*")
do
    protoc -I . -I /usr/local/include \
        --go_out pb \
        --go_opt paths=source_relative \
        --go-grpc_out pb \
        --go-grpc_opt paths=source_relative \
        "${f}"
done

for f in $(find internal -name "*.proto")
do
    protoc -I . -I /usr/local/include \
        --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
        "${f}"
done
