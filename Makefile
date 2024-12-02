

postgres:
	docker run --name wallet -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it wallet createdb --username=root --owner=root wallet

migrate:
	migrate create -ext sql -dir migrations -seq $(name)
	
migrateup:
	migrate -path migrations -database 'postgresql://root:secret@localhost:5432/wallet?sslmode=disable' -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/wallet?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run ./cmd/app


.PHONY: server postgres createdb  migrateup migratedown sqlc
