#!/bin/bash

source update.sh

pushd `go env GOROOT`
GO_VERSION=`git tag -l --contains HEAD`
if [[ -z $GO_VERSION ]]; then
   GO_VERSION=`git rev-parse --verify --short HEAD`
fi
popd

echo $GO_VERSION

pushd "$GOPATH/src/github.com/petar/GoLLRB"
LLRB_VERSION=`git tag -l --contains HEAD`
if [[ -z $LLRB_VERSION ]]; then
   LLRB_VERSION=`git rev-parse --verify --short HEAD`
fi
popd

echo $LLRB_VERSION

OWN_VERSION=`git tag -l --contains HEAD`
if [[ -z $OWN_VERSION ]]; then
   OWN_VERSION=`git rev-parse --verify --short HEAD`
fi

OTHER_FILENAME="_bench/other_heap-"$GO_VERSION"_llrb-"$LLRB_VERSION".bench"
OWN_FILENAME="_bench/own_$OWN_VERSION.bench"


go test ./benchmarks/_bench/. -test.bench . -benchmem -tags other > $OTHER_FILENAME
go test ./benchmarks/_bench/. -test.bench . -benchmem -tags own > $OWN_FILENAME

benchcmp $OTHER_FILENAME $OWN_FILENAME > "bench/$OWN_VERSION_vs_heap"$GO_VERSION"_llrb-"$LLRB_VERSION"
