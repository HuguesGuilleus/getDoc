#!/bin/bash

clear

# Style
cat style/*.css | tr '\t\n' ' ' |
	# sed -E 's#[^/]\*[^/]##g' | sed -E 's#/\*[^*]*\*/# #g' |
	sed -E 's# +# #g' | sed -E 's#([:;{}]) #\1#g' |
	sed 's#^#\t<style>#' | sed 's#$#</style>\n#' > style.html

# Template
cd doc/data/
printf "package data
var index = \`" > index.go
sed '/<!--CODE-->/r ../../style.html' < index.gohtml |
	sed '/<!--CODE-->/d' >>index.go
echo '`' >> index.go
cd ../..

rm style.html

# Build
go build
