# Golang Product Reviews

Sergej Sizov

# Components

Services:
- **API** is a service exposing REST API for managing products and reviews. 
- **Audit** is a service that listens for review changes and logs them.

Infra:
- **PostgreSQL** for persistence.
- **NATS** for messaging.
- **Redis** for caching.

The whole environment is deployed using **Docker Compose**.

# Buiding and running

### Building locally

```shell
make
```

### Building with Docker

```shell
make build-docker
```

### Running with Docker Compose

```shell
docker compose up
```

# Examples

Check `docs` for request examples.

# Design consideration

### Building

There is a make file that automates project build, lint, test.

### Deployment

The application is dockerized, we can easily run it in Docker Compose.

### Rest API

In order to expose a REST API we need to implement an HTTP server. 
I am using Gorilla Mux for request routing.

### Persistence

Products and reviews are stored in Postgres database.
Gorm is used for mapping relational data into structs.

### DB Migrations

Database migrations can be executed on application start.
This is controlled by ENV variable.
Once going into production we would run migration either manually or
on canary pod only.

### Messaging

For purpose of this exercise I am using NATS as it is lightweight
and works out of the box.

### Caching

Average rating and reviews are cached.
Caching is implemented using Redis.
Redis locks are used to implement single flight pattern on cache miss.

### Test coverage

HTTP handlers and service layer are covered with unit tests.
I am using both stubs and mocks where it makes more sense.

### Things to improve

- Tests for Postgres using TestContainers.
- Integration E2E tests.
- Run tests in CI on GitHub Actions.
