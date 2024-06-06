#!/bin/sh

set -e # script exit immidiately if the command returns non-0 status

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@" # take all parameters passed to the script and run it
