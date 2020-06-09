#!/usr/bin/env bash

protoc -I/usr/local/include -I. \
  -I${GOPATH}/src \
  --go_out=plugins=grpc:. \
  *.proto