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
git push origin ft/docker; #copy remote url then paste to browser, click Create pull request, then you'll see Workflow running



```
### test to run ci
##### Github actions postgres: https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers

##### Golang migrate:  https://github.com/golang-migrate/migrate, CLI, CLI Documentation

##### Web Framework Golan Gin: https://github.com/gin-gonic/gin

##### Install  https://github.com/spf13/viper

