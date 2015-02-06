#!/bin/bash

set -e

echo "!! Testing datastructure implementations"
go test -cover ./...

echo "!! Updating datagen templates"
pushd cmd/datagen/ && go generate
popd

echo "!! Verifying code generated for sorted map"
for i in "int" "float64" "string" "[]byte" "[]string"; do
    echo " -key=$i -val=$i"

    go run cmd/datagen/*.go smap -key=$i -val=$i > gen_smap.go 2>/dev/null
    go build gen_smap.go || rm gen_smap.go
    go vet gen_smap.go || rm gen_smap.go
    golint gen_smap.go || rm gen_smap.go
    rm gen_smap.go
done

echo "!! Verifying code generated for sorted set"
for i in "int" "float64" "string" "[]byte" "[]string"; do
    echo " -key=$i"
    go run cmd/datagen/*.go sset -key=$i > gen_sset.go 2>/dev/null
    go build gen_sset.go || rm gen_sset.go
    go vet gen_sset.go || rm gen_sset.go
    golint gen_sset.go || rm gen_sset.go
    rm gen_sset.go
done

echo "!! Verifying code generated for heap"
for i in "int" "float64" "string" "[]byte" "[]string"; do
    echo " -key=$i"
    go run cmd/datagen/*.go heap -key=$i > gen_heap.go 2>/dev/null
    go build gen_heap.go || rm gen_heap.go
    go vet gen_heap.go || rm gen_heap.go
    golint gen_heap.go || rm gen_heap.go
    rm gen_heap.go
done

pushd codegen
echo "!! Generating benchmarked sorted maps"
go run ../cmd/datagen/*.go smap -key string  -val string > smap_string_string.go
go run ../cmd/datagen/*.go smap -key []byte  -val string > smap_bytes_string.go 2> /dev/null
go run ../cmd/datagen/*.go smap -key int     -val string > smap_int_string.go
go run ../cmd/datagen/*.go smap -key float64 -val string > smap_float_string.go

echo "!! Generating benchmarked sorted sets"
go run ../cmd/datagen/*.go sset -key string  > sset_string.go
go run ../cmd/datagen/*.go sset -key []byte  > sset_bytes.go 2> /dev/null
go run ../cmd/datagen/*.go sset -key int     > sset_int.go
go run ../cmd/datagen/*.go sset -key float64 > sset_float.go

echo "!! Generating benchmarked heaps"
go run ../cmd/datagen/*.go heap -key string  > heap_string.go
go run ../cmd/datagen/*.go heap -key []byte  > heap_bytes.go 2> /dev/null
go run ../cmd/datagen/*.go heap -key int     > heap_int.go
go run ../cmd/datagen/*.go heap -key float64 > heap_float.go

echo "!! Check benchmarked types build together"
go build . && go clean
popd


echo "!! All good!"
