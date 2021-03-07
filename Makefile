run-pg:
	docker run --name coins-postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_DB=wallet_db \
		-d -p 5432:5432 postgres:12

run-pg-migrate:
	migrate -source file://migrations/ -database postgres://postgres:postgres@localhost:5432/wallet_db?sslmode=disable up

kill-pg:
	docker stop coins-postgres
	docker rm coins-postgres