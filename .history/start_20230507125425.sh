#!/bin/sh

# Start the first process
set -e

echo "run db migration"
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "start app"
/app/app