#!/bin/bash

clear

# Template
cd doc/data/
printf "package data
var index = \`" > index.go
cat index.gohtml >> index.go
echo '`' >> index.go
cd ../..

# Build
go build
