#!/bin/sh

for f in $(find . -name "*.proto" -not -path "*internal*")
do
    protoc -I . -I /usr/local/include \
        --go_out=plugins=grpc:pb \
        --go_opt=paths=source_relative \
        "${f}"
done

for f in $(find internal -name "*.proto")
do
    protoc -I . -I /usr/local/include \
        --go_out=plugins=grpc:. \
        --go_opt=paths=source_relative \
        "${f}"
done
