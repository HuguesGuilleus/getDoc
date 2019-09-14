---
title: Index
---

This program get the documentation from source code.

## Index
{% include index_file.liquid %}

## Example
This is a very simple example:
```c
#include <stdio.h>

// La fameuse fonction main
int main(int argc, char const *argv[]) {
	printf("Hello World\n");
	return 0;
}
```

To extract the documentation:
```
$ ./getDoc
```

A `doc.html` file is generated, you can see it in your favorite browser or upload it in a web server.

## Building
The HTML template is keep into the binary, and the template include CSS an Js so it' required little manipulation. All this operation are made by `build.bash` script at the root of the project.

Binary files are available: [https://github.com/HuguesGuilleus/getDoc/releases](https://github.com/HuguesGuilleus/getDoc/releases)
