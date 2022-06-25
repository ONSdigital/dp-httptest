#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-mocking
  make audit
popd
