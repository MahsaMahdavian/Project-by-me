include .env
export

MIGRATE_DIR=database/migrations
Db_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable


migrate_up:
	migrate -path $(MIGRATE_DIR) -database $(Db_URL) up

create_migration:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(NAME)

migrate_down:
	migrate -path $(MIGRATE_DIR) -database $(Db_URL) down 

migrate_status:
	migrate -path $(MIGRATE_DIR) -database $(Db_URL) version

migrate_clean:
	migrate -path $(MIGRATE_DIR) -database $(Db_URL) force 2
run:
	go run cmd/main.go

