#!/bin/bash

ENV=${1:-development}

case $ENV in
  development)
    DRIVER="postgres"
    DB_STRING="host=localhost port=5432 user=vyacheslavchumakov password=1111 dbname=poe_chat_api sslmode=disable"
    ;;
  production)
    DRIVER="postgres"
    DB_STRING="host=prod_host port=5432 user=vyacheslavchumakov password=1111 dbname=poe_chat_api sslmode=require"
    ;;
  *)
    echo "Unknown environment: $ENV"
    exit 1
    ;;
esac

goose -dir ./db/migrations "$DRIVER" "$DB_STRING" "${@:2}"
