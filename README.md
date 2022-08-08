<p align="center">
<h1 align="center">Form3 Api Client</h1>
<p align="center">Simple HTTP Api client library made in go allowing interaction with Form3 Accounts Api</p>

## Candidate

Daniel Magalhães

_I'm new to go._

## Prerequisites

The following list of software is based on the versions I've used to build this challenge

- [Go 1.17](https://go.dev/doc/go1.17)
- [Docker 20.10.12](https://docs.docker.com/engine/release-notes/#201012)
- [Docker compose 1.29.2](https://docs.docker.com/compose/release-notes/#1292)
- [GNU Make 4.2.1](https://lists.gnu.org/archive/html/info-gnu/2016-06/msg00005.html)

## Requirements

As described in the submission guide, the solution meets the following requirements:

- Be written in Go.
- Use the `docker-compose.yaml` of this repository.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.

## Project Structure

```tree

├── Dockerfile
├── integration-tests-entrypoint.sh.go
├── wait-for.sh
├── Makefile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── accounts
│   │     ├── accounts.go
│   │     └── accounts_test.go
│   ├── core
│   │     ├── base_client_test.go
│   │     ├── base_client.go
│   │     ├── error_test.go
│   │     ├── error.go
│   │     ├── request_builder_test.go
│   │     ├── request_builder.go
│   │     ├── request_test.go
│   │     ├── request.go
│   │     ├── response_test.go
│   │     └── response.go
│   ├── models
│   │     └── models.go
│   ├── form3.go
│   └── form3_test.go
├── scripts
│   └── db
│       └── 10-init.sql
└── tests
    └── integration
        ├── accounts_test.go
        └── test_utils_test.go
```

## Client Architecture

The Library code related to the Api Client implementation that's ok to use by external applications is inside the `/pgk` directory.

Inside that directory there are 4 go packages defined:

### form3

This is the entrypoint for the Api client usage. It allows the creation of a Form3 Client with possibility to define some options. This is accomplished implementing the options pattern.
It has specific client implementations that are built during the client creation. E.g: Accounts

### core

Contains all the core logic to deal with http requests and responses.
This package has no dependencies to other packages in the project. It should be isolated and generic enough to be used as a base client that can be re-used across specific api calls.

### accounts

This package contains the specific implementation of the Accounts Client.
Since go doesn't support inheritance, composition is being used to achieve the same purpose. This specific client implementation has a base_client that is responsible to make the http requests and handle http responses.
Other specific clients can be created the same way according to the API entities available.

### models

Contains the declaration of the Accounts Api models

## Usage

Complete examples can be found under the `/examples` directory, but here is a brief explanation on how to use the client.

### Fetch

```go

client, _ := form3.NewClient(
  form3.WithBaseUrl(*u),
  form3.WithTimeoutInMilliseconds(30),
)

accountFetchResponse, err := client.Accounts.Fetch(ctx, uuid)

```

### Create

```go

client, _ := form3.NewClient(
  form3.WithBaseUrl(*u),
  form3.WithTimeoutInMilliseconds(30),
)

accountCreationResponse, err := client.Accounts.Create(ctx, newAccount)

```

### Delete

```go

client, _ := form3.NewClient(
  form3.WithBaseUrl(*u),
  form3.WithTimeoutInMilliseconds(30),
)

err = client.Accounts.Delete(ctx, uuid, version)

if err != nil {
    log.Fatalf("Fatal error: %s", err)
}

```

## Tests

Unit and integration tests run when the `docker-compose up` command is executed.
However, there is also a Makefile available to run unit and integration tests by running the following commands:

```bash
make unit-tests
make integration-tests
```
