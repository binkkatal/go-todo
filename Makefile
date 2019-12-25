# makes a migration file

test:
	-dropdb $(DB_NAME)
	createdb $(DB_NAME)
	migrate -path=./pkg/postgres/migrations -database="postgres://localhost/$(DB_NAME)?sslmode=disable" up
	DB_URL=postgres://localhost/$(DB_NAME)?sslmode=disable go test ./...

migration:
	migrate create -ext "sql" -dir "./pkg/postgres/migrations" $(name)

# runs the migrations
migrate:
	migrate -path=./pkg/postgres/migrations -database=$(DB_URL) up

# runs the migrations down
migratedown:
	migrate -path=./pkg/postgres/migrations -database=$(DB_URL) down 1

# runs the migration drop
migratedrop:
	migrate -path=./pkg/postgres/migrations -database=$(DB_URL) drop

# runs the migration version
migrateversion:
	migrate -path=./pkg/postgres/migrations -database=$(DB_URL) version

build:
	go build -o ./bin/http -i ./cmd/http