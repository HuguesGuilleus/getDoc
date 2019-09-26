---
title: Languages Supportés
---

## Note d'introduction
`getDoc` découpe chaque ligne, lui attribut un type grâce à des expression régulière puis les sélectionne. Cela nécessite un minimum d'indentation.

De plus ce programme se limite au objets globaux.


## Commentaire
Ce système de documentation vient du langage Go. La documentation est écrit dans les commentaires juste avant l'élément à documenter. Si on souhaite écrire plusieurs paragraphes, il faut passer une ligne commentée. Un paragraphe peut être découpé sur plusieurs lignes conjointes.
```c
// Un premier paragraphe
//
// Cette macro constante est juste pour l'expemple. Vous ne devez
// pas l'utiliser pour de vrai les enfants.
#define YOLO "yolo!"
```


## Bash
### Commentaire supporté pour la documentation
Un ou plusieurs croisillons: `#` (Le *shebang* est ignoré)

### Pris en charge:
- Fonction (globale) (`func`)
- Variable et constante globale


## C
### Commentaire supporté pour la documentation
Deux barres obliques ou plus: `//`

### Pris en charge:
- Fonction (globale) (`func`)
- Type (`type`)
- Les macros fonctions (`macroFunc`) et variables (`macroConst`)
- Variable globale

## Go
### Commentaire supporté pour la documentation
Deux barres obliques ou plus: `//`

### Pris en charge:
- Fonction (globale) (`func`)
- Type (`type`)
- Variable et constante globale


## Javascript
### Commentaire supporté pour la documentation
Deux barres obliques ou plus: `//`

### Pris en charge:
- Fonction (globale) (`func`)
- Classe (`class`)
- Variable et constante globale
