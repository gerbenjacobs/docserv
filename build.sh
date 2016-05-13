#!/usr/bin/env bash

#formatting
gofmt -w .

#generate html as bindata
go-bindata -pkg bindata -o ./bindata/bindata.go ./resources/template.html

#run the "test"
if [[ $1 = "run" ]]; then
    go-bindata -pkg static -o ./test/static/static.go README.md test/TEST.md
    go run ./test/main.go
fi