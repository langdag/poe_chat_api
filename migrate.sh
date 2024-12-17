#!/bin/bash

# Specify the full path to the YAML file
FILE="./settings/database.yml"

# Ensure the file exists
if [ ! -f "$FILE" ]; then
    echo "Error: File '$FILE' does not exist."
    exit 1
fi

# Parse YAML
DBNAME=$(grep 'dbname:' "$FILE" | awk '{print $2}')
USER=$(grep 'user:' "$FILE" | awk '{print $2}')
PASSWORD=$(grep 'password:' "$FILE" | awk '{print $2}')

ENV=${1:-development}

case $ENV in
  development)
    DRIVER="postgres"
    DB_STRING="host=localhost port=5432 user=$USER password=$PASSWORD dbname=$DBNAME sslmode=disable"
    ;;
  production)
    DRIVER="postgres"
    DB_STRING="host=prod_host port=5432 user=$USER password=$PASSWORD dbname=$DBNAME sslmode=require"
    ;;
  *)
    echo "Unknown environment: $ENV"
    exit 1
    ;;
esac

# Pass the parsed variables to goose
goose -dir ./db/migrations "$DRIVER" "$DB_STRING" "${@:2}"
