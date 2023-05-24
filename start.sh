#!/bin/sh

set -e 

echo "run db migrations" 
/app/migrate -path /app/migration -database "$DB_URL" -verbose up

echo "start the app" 

exec "$@"