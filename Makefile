postgres_run:
	docker run --name e-wallet-app-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456789 -d postgres:14.5-alpine

create_db:
	docker exec -it e-wallet-app-postgres createdb --username=root --owner=root e_wallet_app_v1

drop_db:
	docker exec -it e-wallet-app-postgres dropdb e_wallet_app_v1

migrate_up:
	migrate -path db/migration -database "postgresql://root:123456789@localhost:5432/e_wallet_app_v1?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:123456789@localhost:5432/e_wallet_app_v1?sslmode=disable" -verbose down

sqlc_generate:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

sqlc_compile:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc compile

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc_generate sqlc_compile test server