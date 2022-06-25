#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-mocking
  make build
popd
