---
title: Supported languages
---

`getDoc` split every line, give a type with regular expression and select it. It's required a minimum of indentation.

Furthermore, this program is limited to global object.


## Comment
This documentation system provides the go language. The documentation is in comments just before the documented element. If we want write many paragraph, you must write a commented line. A paragraph can be split on several line.
```c
// A first paragraph.
//
// This macro constant is just for example. You should no use
// in reel project. Isn't it children?
#define YOLO "yolo!"
```


## C
### Comment supported for documentation
More than two slash are also supported: `//`

### Supported elements:
- Function (global) (`func`)
- Type (`type`)
- The macro function (`macroFunc`) and constant (`macroConst`)

### No supported elements:
- global variable


## Go
## Javascript
## Bash
