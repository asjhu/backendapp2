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
git checkout -b ft/docker
git add . 
git commit -m "updated readme"
git push origin ft/docker

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

