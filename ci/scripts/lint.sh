#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-mocking
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
  make lint
popd
