#!/bin/bash

clear

# Template
cd data/
printf "package data
var index = \`" > index.go
cat index.gohtml >> index.go
echo '`' >> index.go
cd ..

# Build
go build
