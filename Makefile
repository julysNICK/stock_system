DB_URL=postgresql://root:secret@localhost:5432/stock_system?sslmode=disable
postgres:
	docker run --name postgres_stock_system -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

createdb:
	docker exec -it postgres_stock_system createdb --username=root --owner=root stock_system

dropdb:
	docker exec -it postgres_stock_system dropdb stock_system

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

initdocker:
	sudo systemctl start docker && sudo docker start postgres_stock_system

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown new_migration initdocker sqlc