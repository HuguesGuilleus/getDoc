---
title: Arguments of command line
---

To get the documentation of specific files, you need to pass it in argument of the program. It's can be a regular file or a directory. By default is the current directory.

Example:
```bash
# Documents specific files
$ getDoc hello.c hello.h
# Documents all the directory
$ getDoc sources/
```

## Output
To set the HTML output, use `-o` or `--output` options and the name of the file or the directory who will contains the file. By default, is the `doc.html` file. If JSON or XML output is enable, the HTML output is disable.

To generate JSON, use `-j` or `--json` option and the name of the file or directory.

To generate XML, use `-x` or `--xml` option and the name of the file or directory.

*Note:* We can pass several output format and file for each of them.


## General flags

| Short Flags | Long Flags  | Description                        |
| :---------- | :---------- | :--------------------------------- |
| `-l`        | `--lang`    | List supported languages and exit. |
| `-h`        | `--help`    | Print the help and exit.           |
| `-v`        | `--verbose` | Enable verbose mode.               |
|             | `--version` | Print program version and exit.    |
