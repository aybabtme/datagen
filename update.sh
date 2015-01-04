#!/bin/bash

set -e

echo "!! Testing datastructure implementations"
go test ./...

echo "!! Updating datagen templates"
cd cmd/datagen/ && go generate && cd ../../

for i in "int" "float64" "string" "[]byte"; do
    echo "!! Verifying code generated for sorted map of $i to $i"

    go run cmd/datagen/*.go smap -key=$i -val=$i > gen_smap.go
    go build gen_smap.go && echo "- it builds!"
    go vet gen_smap.go && echo "- it's vetted!"
    golint gen_smap.go && echo "- it's linted!"
    rm gen_smap.go
done

echo "!! All good!"
