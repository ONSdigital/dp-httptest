#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-mocking
  make test
popd
