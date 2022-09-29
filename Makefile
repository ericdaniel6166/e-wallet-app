postgres:
	docker run --name e-wallet-app-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456789 -d postgres:14.5-alpine

createdb:
	docker exec -it e-wallet-app-postgres createdb --username=root --owner=root e_wallet_app_v1

dropdb:
	docker exec -it e-wallet-app-postgres dropdb e_wallet_app_v1

migrateup:
	migrate -path db/migration -database "postgresql://root:123456789@localhost:5432/e_wallet_app_v1?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123456789@localhost:5432/e_wallet_app_v1?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test