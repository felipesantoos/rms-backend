#!/bin/bash

echo "\
\033[93m
+----------------------------------+
| Loading environment variables... |
+----------------------------------+\
\033[0m
"
source ./src/ui/api/app/.env-for-local
schema=$(echo $DATABASE_SCHEMA | sed "s/\r//")
user=$(echo $DATABASE_USER | sed "s/\r//")
password=$(echo $DATABASE_PASSWORD | sed "s/\r//")
host=$(echo $DATABASE_HOST | sed "s/\r//")
port=$(echo $DATABASE_PORT | sed "s/\r//")
name=$(echo $DATABASE_NAME | sed "s/\r//")
sslmode=$(echo $DATABASE_SSL_MODE | sed "s/\r//")
migrations_path=$(echo $DATABASE_MIGRATIONS_PATH | sed "s/\r//")
uri="$schema://$user:$password@$host:$port/$name?sslmode=$sslmode"

if [[ schema != "" && migrations_path != "" ]]; then
  echo "Searching for migrations..."
  migrate -source file://$migrations_path -database $uri up
fi
