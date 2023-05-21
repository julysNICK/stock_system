#!/bin/bash

DB_URL="postgresql://root:secret@localhost:5432/stock_system?sslmode=disable"

SQL_QUERY="SELECT * FROM users"

RESULT=$(psql $DB_URL -c "$SQL_QUERY" -t)

if [$? -ne 0]; then 
    echo "Error: $RESULT"
    exit 1
fi

echo "Success: $RESULT"
exit 0
