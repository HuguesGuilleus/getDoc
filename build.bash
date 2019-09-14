#!/bin/bash

# getDoc
# 2019 GUILLEUS Hugues <ghugues@netc.fr>
# BSD 3-Clause "New" or "Revised" License

JS_FILES=search/*.js
JS_NOLETTER=':;,.{}()<>|&=+-'

clear

# Style files
cat style/*.css | tr '\t\n' ' ' |
	# sed -E 's#[^/]\*[^/]##g' | sed -E 's#/\*[^*]*\*/# #g' |
	sed -E 's# +# #g' | sed -E 's#([:;{}]) #\1#g' |
	sed 's#^#\t<style>#' | sed 's#$#</style>\n#' > include.html

# JavaScript files
for file in $JS_FILES
do
	sed 's#//.*##' $file
done | tr '\n\t' '  ' | sed -E 's# +# #g' |
	sed -E "s#([$JS_NOLETTER]) #\1#g" | sed -E "s# ([$JS_NOLETTER])#\1#g" |
	sed 's#{{#{ {#g' | sed 's#}}#} }#g' |
	sed 's#^#\t<script defer>#' | sed 's#$#</script>\n#' >> include.html

# Template
cd doc/data/
printf "package data
var index = \`" > index.go
sed '/<!--CODE-->/r ../../include.html' < index.gohtml |
	sed '/<!--CODE-->/d' >>index.go
echo '`' >> index.go
cd ../..

rm include.html

# Build
go build
