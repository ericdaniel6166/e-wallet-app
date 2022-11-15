DB_URL=postgresql://root:123456789@localhost:5432/e_wallet_app_v1?sslmode=disable

postgres_run:
	docker run --name e-wallet-app-postgres --network e-wallet-app-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456789 -d postgres:14.5-alpine

create_db:
	docker exec -it e-wallet-app-postgres createdb --username=root --owner=root e_wallet_app_v1

drop_db:
	docker exec -it e-wallet-app-postgres dropdb e_wallet_app_v1

migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

aws_migrate_up:
	migrate -path db/migration -database "postgresql://root:XKhsqsLnVoSO4q2ix2Bn@e-wallet-app.cudifdjlfrok.us-east-1.rds.amazonaws.com:5432/e_wallet_app_v1" -verbose up

migrate_up_1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrate_down_1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

migrate_create:
	migrate -path db/migration -database "$(DB_URL)" create -ext sql -dir db/migration -seq init_schema

sqlc_generate:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

sqlc_compile:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc compile

test:
	go test -v -cover ./...

server:
	go run main.go

docker_build:
	docker build -t e-wallet-app:latest .

docker_network:
	docker network create e-wallet-app-network

docker_run:
	docker run --name e-wallet-app-v1 --network e-wallet-app-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:123456789@e-wallet-app-postgres:5432/e_wallet_app_v1?sslmode=disable" e-wallet-app:latest

proto_generate_win:
	del /f protogen\*.go
	protoc --proto_path=proto --go_out=protogen --go_opt=paths=source_relative \
	--go-grpc_out=protogen --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=protogen --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

proto_generate:
	rm -f protogen/*.go
	protoc --proto_path=proto --go_out=protogen --go_opt=paths=source_relative \
	--go-grpc_out=protogen --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=protogen --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 8081 -r repl

.PHONY: evans proto_generate_win proto_generate postgres_run create_db drop_db migrate_up migrate_up_1 migrate_down migrate_down_1 migrate_create sqlc_generate sqlc_compile test server docker_build aws_migrate_up