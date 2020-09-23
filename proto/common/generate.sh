#!/usr/bin/env bash

# Clean up the current generated protobuf go codes
find . -name \*.pb.go -delete

GOSRC="$GOPATH"/src

protoc *.proto \
       --go_out=plugins=grpc:"$GOSRC" \
       --grpc-gateway_out=logtostderr=true:"$GOSRC" \
       -I=. \
       -I="$GOSRC" \
       -I="$GOSRC/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"

