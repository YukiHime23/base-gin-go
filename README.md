# Base template gin-go
Golang base template for API server

## Setup
- Install packages:
```
go mod install
```
- Copy `.env` from `.env.example`:  
```
cp .env.example .env
``` 
- Update `.env`

## Rules
- `models`: entities defination
- `controllers`: binding data, validation
- `services`: write business logic
- `resositories`: get data from storage (DB, firebase, bigquery,..)
- `config`: setup or read configuration variables
- `pkg`: internal package

## Unit test
- Setup gomock:
```
go install github.com/golang/mock/mockgen@v1.6.0
```
- Generate mock repository:
```
mockgen -source repositories/user.go -package mock -destination repositories/mock/user.go
```

- Setup ginkgo:
```
go get github.com/onsi/ginkgo/v2/ginkgo
```
```
go get github.com/onsi/gomega/...
```
- Create test suite:
```
cd to/pkg/is/tested
ginkgo bootstrap
ginkgo generate %FILENAME%
```
- Run test:
```
ginkgo ./usecase
// or
go test ./usecase
// or run all test in project
ginkgo ./...
```

## DB Migration

- Install cmd package:
  ```
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```
- Create migration files:
  ```
  migrate create -ext sql -dir db/migration -seq <file_name>
  ```
- Run migrate:
  - UP:
    ```
    migrate -path db/migration -database "postgres://<user>:<pwd>@localhost:5432/<db_name>?sslmode=disable" -verbose up
    ```
  - DOWN:
    ```
    migrate -path db/migration -database "postgres://<user>:<pwd>@localhost:5432/<db_name>?sslmode=disable" -verbose down
    ```
  - Or use makefile
    ```
    $ make migrateup step=1
    $ make migratedown step=2
    ```

## Used
- gin-gonic
- air
- gorm
- crypto/bcrypt
- mockgen 1.6.0
- jwt v4
- golang-migrate
