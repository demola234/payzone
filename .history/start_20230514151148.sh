#!/bin/sh

# Start the first process
set -e

echo "run db migration"
source /app/.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start app"
exec "$@"