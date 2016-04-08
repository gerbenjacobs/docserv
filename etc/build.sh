#!/usr/bin/env bash
if [[ $PWD = */etc ]]; then
    echo "Run this from the root of the project"
    exit 1
fi

gofmt -w .
go-bindata -pkg docserv -o ./bindata.go ./etc/doc.html

if [[ $1 = "run" ]]; then
    go run ./test/main.go
fi