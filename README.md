# golang-product-reviews
Golang Product Reviews

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

# Design consideration

In order to expose a REST API we need to implement an HTTP server. 
I am using Gorilla Mux for request routing.

Products and reviews are stored in Postgres database.
Database migrations will be applied on application start.
Gorm is used for mapping relational data into structs. 

