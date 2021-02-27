#!/bin/sh

# install dependencies
go mod tidy && go install \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc


# remove old pb files
find ./pb -name "*.pb.go" -delete


# generate protobuf file
for f in $(find ./proto/daangn -name "*.proto")
do
  protoc -I ./proto \
    --go_out pb \
    --go_opt paths=source_relative \
    --go-grpc_out pb \
    --go-grpc_opt paths=source_relative \
    "${f}"
done


# move and delete directories
cp -fR pb/daangn/eboolkiq/* pb/ \
&& rm -rf pb/daangn
