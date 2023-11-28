``` sh
Clone with ssh recommended
$ git clone git@github.com:AndresXLP/reto-clubhub.git

Clone with https
$ git clone https://github.com/AndresXLP/reto-clubhub.git
```

# Requirements

* go v1.20
* go modules

# Technology Stack

- [echo](https://echo.labstack.com/)
- [validator](https://github.com/go-playground/validator)
- [GORM](https://gorm.io/)

# Build

* Install dependencies:

```sh
$ go mod download
```

* [Migrations](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 
```sh
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
$ migrate -path ./internal/infra/resource/postgres/migrations -database postgresql://${POSTGRESQL_DB_USER}:${POSTGRESQL_DB_PASSWORD}@${POSTGRESQL_DB_HOST}:${POSTGRESQL_DB_PORT}/${POSTGRESQL_DB_NAME}?sslmode=disable up
```

* Run local
```sh
$ go run cmd/main.go
```

* Run with Docker:

```sh 
$ make compose-up 
```

# Environments

#### Required environment variables

* `SERVER_PORT`: port for the server
* `POSTGRES_HOST`: host database
* `POSTGRES_USER`: user database
* `POSTGRES_PASSWORD`: password database
* `POSTGRES_NAME`: name database
* `POSTGRES_PORT`: port database


# Contributors

* Andres Puello

