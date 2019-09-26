---
title: Supported languages
---

`getDoc` split every lines, give a type with regular expressions and select it. It's required a minimum of indentation.

Furthermore, this program is limited to global object.


## Comment
This documentation system provides from the go language. The documentation is in comments just before the documented element. If we want write many paragraph, you must write a commented line. A paragraph can be split on several line.
```c
// A first paragraph.
//
// This macro constant is just for example. You should no use
// in reel project. Isn't it children?
#define YOLO "yolo!"
```


## Bash
### Comment supported for documentation
One or more hash: `#` (The shebang are ignored)

### Supported element:
- Function (global) (`func`)
- Global Variable


## C
### Comment supported for documentation
More than two slash are also supported: `//`

### Supported element:
- Function (global) (`func`)
- Type (`type`)
- The macro function (`macroFunc`) and constant (`macroConst`)
- Global variable


## Go
### Comment supported for documentation
More than two slash are also supported: `//`

### Supported element:
- Function (global) (`func`)
- Type (`type`)
- Global variable and constant


## Javascript
### Comment supported for documentation
More than two slash are also supported: `//`

### Supported element:
- Function (global) (`func`)
- Class (`class`)
- Global variable and constant
