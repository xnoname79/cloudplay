# cloudplay

A minimal Go web server that will evolve into a cloud gaming service capable of running PS2 games. For now it exposes a few endpoints and serves as the project's starting point.

## Building

```bash
go build
```

## Running

```bash
PORT=8080 go run .
```

The server listens on the port specified by the `PORT` environment variable (defaults to `8080`). It currently provides the following endpoints:

- `/` – returns **Hello, world!**
- `/health` – returns a JSON health status
- `/session/start` – placeholder for starting a PS2 game session
- `/games` – lists available games

The `/session/start` endpoint now requires a `game` query parameter indicating
which game to start.

These endpoints will expand as the project grows toward streaming PS2 games through the cloud.
