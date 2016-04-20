#!/usr/bin/env bash
if [[ $PWD = */resources ]]; then
    echo "Run this from the root of the project"
    exit 1
fi

#formatting
gofmt -w .

#generate html as bindata
go-bindata -pkg bindata -o ./bindata/bindata.go ./resources/doc.html

#run the "test"
if [[ $1 = "run" ]]; then
    go run ./test/main.go
fi