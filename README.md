# User Storage

## Execution

### Environment variables

* DATABASE_PATH - sqlite3 database file path
* PORT - server port

### Migrations

1. install: https://github.com/golang-migrate/migrate
2. run: 
```shell
migrate -source file://${PWD}/migrations -database sqlite3://${DATABASE_PATH} up
```

## Develop

### RESTAPI generation

1. install: https://github.com/go-swagger/go-swagger
2. run:
```shell
swagger generate server -t gen -f swagger.yaml --exclude-main -A user_storage
```
