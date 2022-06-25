#!/bin/bash -eux

cwd=$(pwd)

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

pushd $cwd/dp-mocking
  make lint
popd
