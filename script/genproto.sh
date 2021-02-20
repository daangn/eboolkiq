#!/bin/sh

# install dependencies
go mod tidy && go install \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc


# remove old pb files
rm -rf pb/*


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
mv pb/daangn/eboolkiq/* pb/ \
&& find ./pb -type d -empty -delete
