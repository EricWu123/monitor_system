#!/bin/bash

cd "$(dirname $0)" || exit 1

cd ..

go test -covermode=count -coverprofile=scripts/cover.out -coverpkg=$(go list ./... | grep -v "^monitor_system$" | tr '\n' ',') ./...
go tool cover -html=scripts/cover.out -o scripts/converage.html