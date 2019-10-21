---
title: Développement interne
---

Les fonctions d'analyse sont de deux catégories: celles qui attribuent un type à chaque ligne, et celles qui listent les éléments. Ces fonctions travaillent par paire et sont référencées dans la variable `parserList` dans `doc/parser.go`.

## Test unitaire
```go
func TestLangXxx_parse(t *testing.T) {
	testingFile = "file.ext"
	testParser(t, "NOM", []string{
		// lignes d'entré
	}, Element{
		Name:     "",
		LineName: "",
		Type:     "",
	})
}
```
`testParser` règle des valeurs par détaut:
- `Comment` est définit à `[]string{"Comment"}`.
- Si `LineNum` n'est pas défini, il est définit à `2` et une ligne de commentaire est ajoutée au tout début de la liste de ligne.
- Si `Lang` n'est la défini alors, il est défini comme l'extension (sans le point) du fichier.

```go
func TestLangXxx_type(t *testing.T) {
	testType(t, langXxx_type, []testingLine{
		// {type attendue, ligne d'entré, ligne attendue de sortie}
	})
}
```
