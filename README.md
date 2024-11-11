name: ci-test

on:
  push:
    branches: [ "master" ]
  pull_request:
    branches: [ "master" ]

jobs:

  test:
    name: Test
    runs-on: ubuntu-latest

    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_USER: root
          POSTGRES_PASSWORD: secret
          POSTGRES_DB: app3
        ports:
          - 5432:5432
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: ^1.20
      id: go
    - name: Install golang-migrate
      run: |
        curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz
        sudo mv migrate /usr/bin/
        which migrate
    - name: Run migration
      run: make migrateup

    - name: Test
      run: make test



##### Install migrate https://github.com/golang-migrate/migrate/tree/master > CLI Documentation
##### Installation>Downloads>Linux, extract, move /usr/local/bin: https://github.com/kyleconroy/sqlc
```sh
docker pull postgres:17-alpine
docker images; 
make postrgres
make createdb
make migrateup
go mod init github.com/asjhu/backendapp
go mod tidy


migrate -version
migrate create -ext sql -dir db/migration -seq init_schema
make dropdb
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/backend?sslmode=disable" -verbose up
make migratedown; make migrateup
sqlc version; sqlc init
go get github.com/lib/pq # go get github.com/lib/pq
go get github.com/stretchr/testify  # https://github.com/stretchr/testify
go get -u github.com/gin-gonic/gin
make server
make test # to create more account
go get github.com/spf13/viper
SERVER_ADDRESS=0.0.0.0:8081 make server # to test temp env variable

```
### test to run ci
##### Github actions postgres: https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers

##### Golang migrate:  https://github.com/golang-migrate/migrate, CLI, CLI Documentation

##### Web Framework Golan Gin: https://github.com/gin-gonic/gin

##### Install  https://github.com/spf13/viper

