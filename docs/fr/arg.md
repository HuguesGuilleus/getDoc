---
title: Arguments
---

Pour récupérer la documentation sur des fichiers spécifiques, il suffit de les passer directement arguments. Cela peut-être des fichiers ou des répertoires. Par défaut c'est le répertoire courant.

Exemple:
```bash
# Documenter des fichiers spécifiques
$ getDoc hello.c hello.h
# Documenter tout un répertoire
$ getDoc sources/
```


## Sortie
Pour régler la sortie en HTML, utilisez les options `-o` ou `--output` suivi du nom du fichier ou du répertoire dans lequel enregistrer le fichier. Par défaut c'est le fichier `doc.html`. Si la sortie en JSON ou XML est activée, la sortie par défaut du HTML est désactivée.

Pour générer du JSON, utilisez les options `-j` ou `--json` suivi du nom du fichier ou du répertoire.

Pour générer du XML, utilisez les options `-x` ou `--xml` suivi du nom du fichier ou du répertoire.

Notons que l'on peut indiquer plusieurs formats et plusieurs sorties pour chacun d'eux.


## Drapeaux Généraux

| Drapeau short  | Drapeau long   | Description                                         |
| :------------- | :------------- | :-------------------------------------------------- |
| `-l`           | `--lang`       | Liste les langages supportés et quitte le programme |
| `-h`           | `--help`       | Affiche l'aide et quitte le programme               |
| `-v`           | `--verbose`    | Active le mode verbeux                              |
|                | `--version`    | Affiche la version et quitte le programme           |
