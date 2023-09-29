
migrate-tool-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	migrate create -ext sql -dir ./migrations/postgres -seq $(filename)

migrate-up:
	migrate -database 'postgresql://api-server-user:12345678@localhost:5432/api-server-db?sslmode=disable' -path ./migrations/postgres -verbose up

migrate-down:
	migrate -database 'postgresql://api-server-user:12345678@localhost:5432/api-server-db?sslmode=disable' -path ./migrations/postgres -verbose down

migrate-force:
	migrate -database 'postgresql://api-server-user:12345678@localhost:5432/api-server-db?sslmode=disable' -path ./migrations/postgres force $(version)

.PHONY: migrate-tool-install migrate-create migrate-up migrate-down migrate-force