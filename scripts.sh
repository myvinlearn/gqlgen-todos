#!/bin/bash

go build -o /bin/demo
./bin/demo && go test -v ./...