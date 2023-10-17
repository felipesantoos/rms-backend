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

echo "\
\033[93m
+----------------------------------------------+
| Starting databases (PostgreSQL and Redis)... |
+----------------------------------------------+\
\033[0m
"
docker compose rm -sf
docker compose up database redis --build -d

echo "\
\033[93m
+-------------------------------------+
| Downloading project dependencies... |
+-------------------------------------+\
\033[0m
"
go mod tidy

echo "\
\033[93m
+---------------------------------+
| Generating API documentation... |
+---------------------------------+\
\033[0m
"
bash -c "cd src/ui/api/app && swag init -g ../api.go --output ../docs --dir ../handlers"

echo "\
\033[93m
+---------------------------------------------------------+
| Waiting the database from PostgreSQL to be available... |
+---------------------------------------------------------+\
\033[0m
"
chmod +x ./tools/executables/wait-for-it.sh
./tools/executables/wait-for-it.sh --timeout=20 $DATABASE_HOST:$DATABASE_PORT

echo "\
\033[91m
+-----------------------------------------------------------------------------------------+
| Please check if the database from PostgreSQL is running and run the following commands: |
| $ source tools/executables/load-migrations.sh                                           |
| $ source tools/executables/start-server.sh                                              |
+-----------------------------------------------------------------------------------------+\
\033[0m
"
