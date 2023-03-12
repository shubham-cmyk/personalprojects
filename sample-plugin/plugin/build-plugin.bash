#!/bin/bash

# Build all plugin

go build -buildmode=plugin -o chinese/chinese.so chinese/chinese.go
go build -buildmode=plugin -o english/english.so english/english.go
go build -buildmode=plugin -o russian/russian.so russian/russian.go
