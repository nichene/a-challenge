## Stone-challenge project

API service developed for a challenge.

## Pre configurations

- To load env vars copy the `.env.example` to a `.env` file and modify accordingly.

- Point your DB vars to your desired DB. To run a local DB use the `make db` command.

## Migrations

- It is necessary to install Goose to run DB migrations. Please refer to the docs at: [goose](https://github.com/pressly/goose) for instalation guide.

- To run migrations use `make migrate-up`.

## API

To run the server locally you can type `make run-api`.

## Testing

Run the project tests by running `make tests`.

## Dockerfile

You are able to build the projects image by running:
`build-image`

and check it out by running
`docker run --rm -p 8080:8080 stone-challeng:latest`

## Other Commands

Go through makefile to checkout all available commands for this project.
