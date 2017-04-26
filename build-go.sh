#!/bin/sh

set -eufx

go get -v
GOARCH=arm GOOS=linux GOARM=6 go build -ldflags "-w -s -X main.version=$VERSION -X main.commit=$COMMIT" -o piriscope
