# getDoc

A simple documentation extractor from source Code. It support Bash, C, Js and Go.

## Documentation

-   [Go module reference: ![](https://pkg.go.dev/badge/github.com/HuguesGuilleus/getdoc.svg)](https://pkg.go.dev/github.com/HuguesGuilleus/getDoc@master/doc)
-   [CLI & manual: (English, Français) https://huguesguilleus.github.io/getDoc/](https://huguesguilleus.github.io/getDoc/)

## Build

```bash
git clone https://github.com/HuguesGuilleus/getDoc
cd getDoc
go build
```

You can also use binary: <https://github.com/HuguesGuilleus/getDoc/releases>

## Usage

Run the binary at the root of a project and it will read the file (Bash, C, Go, Js), get the documentation (comments before the element) and write `doc.html` (XML and JSON are also supported).
