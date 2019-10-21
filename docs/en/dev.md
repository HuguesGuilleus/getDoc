---
title: Intern developpement
---

The parser function are divised in two category: they who give a type for each line and they who index the item. The functions work by pair and are indexed in `parserList` variable in `doc/parser.go`.


## Unit test
```go
func TestLangXxx_parse(t *testing.T) {
	testingFile = "file.ext"
	testParser(t, "NAME", []string{
		// input lines
	}, Element{
		Name:     "",
		LineName: "",
		Type:     "",
	})
}
```
`testParser` set default values:
- `Comment` is set to `[]string{"Comment"}`.
- If `LineNum` is not defined, it set to `2` and a comment line is added to the list of line.
- If `Lang` is undefined, it set to extension (without the dot) of the file.

```go
func TestLangXxx_type(t *testing.T) {
	testType(t, langXxx_type, []testingLine{
		// {type expected, input line, output line expected}
	})
}
```
