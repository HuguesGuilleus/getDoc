---
title: Sommaire
---

Ce programme permet d'extraire rapidement la documentation issue d'un code source.

## Sommaire
{% include index_file.liquid %}

## Exemple
Voici un programme très simple:
```c
#include <stdio.h>

// La fameuse fonction main
int main(int argc, char const *argv[]) {
	printf("Hello World\n");
	return 0;
}
```

Pour extraire la documentation:
```
$ ./getDoc
```

Le fichier `doc.html` est généré, vous pouvez directement le visualiser dans votre navigateur préféré ou l'envoyer sur un serveur web.

## Compilation
Le modèle HTML est stocké directement avec le code CSS et HTML dans le binaire. Cela nécessite une préparation.
Tout est effectué dans le fichier `build.bash` à la racine du projet.

Des binaires sont disponibles: [https://github.com/HuguesGuilleus/getDoc/releases](https://github.com/HuguesGuilleus/getDoc/releases)
