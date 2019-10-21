---
title: Interne
---

## Test unitaire

```go
func TestLangXxx_parse(t *testing.T) {
	testingFile = "file.ext"
	testParser(t, "function", []string{
		// lines for test
	}, Element{
		Name:     "",
		LineName: "",
		Type:     "",
	})
}
```

```go
func TestLangXxx_type(t *testing.T) {
	testType(t, langXxx_type, []testingLine{
		{0,"",""},
	})
}
```
