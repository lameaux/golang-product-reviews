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

The application is dockerized, we can easily run it in Docker Compose.

There are 2 services: api server and audit logger.

In order to expose a REST API we need to implement an HTTP server. 
I am using Gorilla Mux for request routing.

HTTP handlers and service layer are covered with unit tests.

Products and reviews are stored in Postgres database.
Database migrations are applied on application start.
Gorm is used for mapping relational data into structs. 

Caching is implemented using Redis.
Messaging is implemented using NATS.


