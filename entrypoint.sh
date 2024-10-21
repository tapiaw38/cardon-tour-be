#!/bin/sh

# Wait for PostgreSQL to be ready
until pg_isready --host=cardon-postgres-db --port=5432 --username=postgres --dbname=cardon-tour-db
do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 2
done

# PostgreSQL is now ready, to run the Go application
echo "========== Starting Go application =========="
exec go run ./cmd/api/ --host 0.0.0.0