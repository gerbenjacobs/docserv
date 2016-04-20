#!/usr/bin/env bash

# set file names
if [[ -z $1 ]]; then
    echo "You need to supply the file paths."
    echo "I.e. ./static.sh README.md docs/CHANGELOG.md"
    exit 1
fi

# get bindata binary
bindata=`which go-bindata`
if [[ $? -ne 0 ]]; then
    echo "You need to install go-bindata. See the README.md for more information"
    exit 1
fi

#generate html as bindata
${bindata} -pkg static -o ./static/static.go $@