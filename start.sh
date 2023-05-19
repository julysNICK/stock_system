#!/bin/sh

set -e 

echo "run db migrations" /app/migrate -path /app/migration -database "postgresql://root:secret@localhost:5432/stock_system?sslmode=disable" -verbose up

echo "start the app" 

exec "$@"