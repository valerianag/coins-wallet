
### API documentation
You can find API documentation with requests examples in [docs/api.md](docs/api.md)

You can find example usage case in [docs/example.md](docs/example.md) (ypu need to run services first - see next section)
### Run service locally with docker-compose:

You can run all needed environment in docker-compose with 
only 2 commands (to run containers and migration):
```
> make dev-up
> make dev-env
```
Then kill it with:
```
> make dev-down
> make dev-clean
``` 
### Run service locally with go run:
```
> make run-pg
> make run-pg-migrate
> go run main.go
```

You can change database DSN with `POSTGRES_DSN` env var like:
```
> export POSTGRES_DSN="postgresql://postgres:postgres@localhost:5432/wallet_db?sslmode=disable"
```

You can change app server port with `APP_PORT` env var like:
```
> export APP_PORT=":8080"
```

### Database
Postgresql used for data storing. 
To run it locally for develop proposes (with docker):

```> make run-pg```

To stop and remove:

```> make kill-pg```

[golang-migrate](https://github.com/golang-migrate/migrate) 
cli used for postgres migrations. You can run it with:

```> make run-pg-migrate``` 

