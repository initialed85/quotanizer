#!/usr/bin/env bash

set -e -x

go test -v ./pkg/file
go test -v ./pkg/quota
go test -v ./pkg/sizer
go test -v ./pkg/sorter
go test -v ./pkg/walker
