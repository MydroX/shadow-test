# Shadow technical test

## Brief

Ta mission si tu l'acceptes, est de faire une API Rest en Golang permettant de gérer les joueurs et équipes d'un club de football club. 

Tu es libre de choisir ton architecture. Penses tout de même faire des packages réutilisables.

Tu auras un `repository` pour enregistrer les données dans une bdd postgres (image docker). Tu es libre d'utiliser un ORM ou pas.

Il s'agira de faire des opérations basiques CRUD sur les `teams` et `players` et d'affecter un joueur à une équipe.

Si tu ne termines pas en quelques heures, ce n'est pas bloquant explique nous pourquoi.

Si tu as du temps et de l'énergie, tu peux faire en plus:
- OpenAPI
- tests unitaires et/ou end-to-end
- Dockerfile
- CI/CD
- Makefile

Pousse ton code sur Github ou Gitlab pour faciliter la review.

## Installation

Create `.env` file with the following content:
```markdown
APP_PORT=3000
APP_ENV="dev"

POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgres"
POSTGRES_DB="shadow"
POSTGRES_HOST="shadow-test-postgres"
POSTGRES_PORT="5432"
```

```bash
docker-compose up --build
```

## API Documentation

| Method | Route | Description | Body |
| --- | --- | --- | --- |
| Teams| | | |
| GET | /teams | Get all teams | - |
| GET | /teams/{id} | Get a team by id | - |
| POST | /teams/create | Create a team | `{"name": "AJ Auxerre", "tag_name": "AJA"}` |
| PUT | /teams/{id}/update | Update a team | `{"name": "AJ Auxerre", "tag_name": "AJA"}` |
| DELETE | /teams/{id}/delete | Delete a team | - |
| Players| | | |
| GET | /players | Get all players | - |
| GET | /players/{id} | Get a player by id | - |
| POST | /players/create | Create a player | `{"surename": "Hein", "firstname": "Gauthier", "dob": "07/08/1996", "number": 7, "position": "midfielder", "nationality": "france", "team_id": 1}` |
| PUT | /players/{id}/update | Update a player |  `{"surename": "Hein", "firstname": "Gauthier", "dob": "07/08/1996", "number": 7, "position": "midfielder", "nationality": "france", "team_id": 1}` |
| DELETE | /players/{id}/delete | Delete a player | - |