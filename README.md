# Quiz de Pâques

Ce projet est un petit quiz de Pâques réalisé en Go et HTML. Il vous permet de jouer à un quiz interactif sur le thème de Pâques.

## Fonctionnalités

- Affichage de questions avec plusieurs options de réponse
- Validation des réponses et affichage du score final
- Possibilité de redémarrer le quiz

## Prérequis

Avant de commencer, assurez-vous d'avoir installé Go sur votre système.

## Installation

1. Clonez ce dépôt sur votre machine :

git clone git@github.com:Dorianlmt/Quiz.go.git

2. Accédez au répertoire du projet :

cd quiz-de-paques

3. Exécutez le fichier quiz.go à l'aide de la commande suivante :

go run quiz.go 
ou
make run

4. Ouvrez votre navigateur et accédez à l'URL suivante :

http://localhost:9999


## Utilisation

- Sélectionnez une réponse pour chaque question en cliquant sur le bouton correspondant.
- Une fois toutes les questions répondues, vous serez redirigé vers une page de fin où vous pourrez voir votre score.
- Vous pouvez redémarrer le quiz à tout moment en cliquant sur le bouton "Recommencer le quiz".

## Structure du projet

- `quiz.go`: Contient le code principal du serveur HTTP et la logique du quiz.
- `src/`: Répertoire contenant les fichiers HTML, CSS et les images utilisées dans le projet.
- `src/templates/`: Contient les templates HTML utilisés par le serveur pour générer les pages.
- `src/img/`: Contient les images utilisées dans le projet.

## Auteur

[Dorian]

N'hésitez pas à contribuer en ouvrant une issue ou une pull request si vous avez des suggestions d'amélioration ou si vous avez rencontré des problèmes.

