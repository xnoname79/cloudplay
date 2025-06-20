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
- `/login` – obtain an access and refresh token
- `/refresh` – refresh an access token using a refresh token
- `/session/start` – start a new game session
- `/session/stop` – stop a running session
- `/session/get?id=ID` – retrieve session metadata

These endpoints will expand as the project grows toward streaming PS2 games through the cloud.

## Video Module

The `video` package now includes an extensible backend interface to hook up a PS2 emulator. `PCSX2Backend` launches a PCSX2 process, while `DummyBackend` is used for tests. Frame capture for PCSX2 is not yet implemented, but the interface provides the foundation for integrating real emulator output.

