DB_URL=postgresql://root:secret@localhost:5432/stock_system?sslmode=disable
network:
	sudo docker network create stock-system-network
clearforup:
	sudo docker rm stock_system-api-1 && sudo docker rm stock_system-postgres-1 && sudo docker rmi stock_system-api

dockerbuild:
	sudo docker build -t stock_system:latest .

dockefilerun:
	sudo docker run --name stock_system --network stock-system-network -p 8080:8080 -e GIN_MODE=release -e DB_URL="postgresql://root:secret@postgres_stock_system:5432/stock_system?sslmode=disable" stock_system:latest 

postgres:
	sudo docker run --name postgres_stock_system --network stock-system-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

createdb:
	sudo docker exec -it postgres_stock_system createdb --username=root --owner=root stock_system

dropdb:
	sudo docker exec -it postgres_stock_system dropdb stock_system

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

initdocker:
	sudo systemctl start docker && sudo docker start postgres_stock_system

composeup:
	sudo docker compose up 
	
sqlc:
	sqlc generate

test:
	go test -v -cover ./...
server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/julysNICK/stock_system/db/sqlc StoreDB
.PHONY: network clearforup postgres createdb dropdb migrateup migratedown migrateup1 new_migration initdocker composeup sqlc test server 