# loan-service

Small scale loan engine for a loan system.

A service that provide API for processing user's loan data.

## Getting Started

```text
Make sure go version >=1.20 is already installed on your machine.
All commands below should be run on project root folder.
```

### Pull dependencies

```shell
go mod tidy && go mod vendor -v
```

### Setup your environment

```shell
make setup
```

### Setup your own credential config

```shell
code files/etc/credential/development/loan-service.secret.json
```

### Run dependencies and the application

```shell
make docker-start
```

## Contributor

* Evin Cintiawan (ecintiawan)