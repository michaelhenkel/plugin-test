#!/bin/bash
go build -buildmode=plugin -o foo/foo.1.so foo/foo.1.go
go build -buildmode=plugin -o bar/bar.1.so bar/bar.1.go
go build
