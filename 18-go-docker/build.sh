#!/bin/bash

find . -name 'main.go' | while read line; do
    echo "Compiling '$(dirname $line)'"
    pushd $(dirname $line) > /dev/null
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
    popd > /dev/null
done

echo "Building containers"

docker-compose build
