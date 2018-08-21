#!/usr/bin/env bash

cp -r $GOROOT/src/crypto/tls/* .
cp $GOROOT/src/internal/cpu/* internal/cpu/
cp $GOROOT/src/internal/testenv/* internal/testenv

# Fixing imports.
for dir in . ./internal/cpu ./internal/testenv; do
  pushd ${dir}
  sed -i 's/\"internal/\"github.com\/gortc\/dtls\/internal/g' *.go
  sed -i 's/\"golang_org/"golang.org/g' *.go
  sed -i 's\package tls\package dtls\g' *.go
  goimports -w -local github.com/gortc/dtls/ .
  popd
done
