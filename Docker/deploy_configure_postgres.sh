#!/bin/bash

#Script to create stocks database and create tables.

touch ./db_password.txt
echo "$1" > ./db_password.txt
if [ -z "$1" ]; then
    echo "Warning: No Administrator password provided."
    echo "Usage: ./deploy_configure_postgres.sh <Administrator Password>"
    exit 1
fi
docker compose --log-level ERROR -f postgres-compose.yaml up
sleep 5
postgres_installed=`docker ps | grep postgres`
if [ -z "$postgres_installed"]; then
    echo "Failed to deploy postgres container"
    exit 1
fi
postgres_container_id=`docker ps | grep postgres | awk '{print$1}'`
echo "$postgres_container_id" 
docker cp ./DB_Create.sql $postgres_container_id:/
docker exec -it $postgres_container_id bash -c 'psql -U postgres -f /DB_Create.sql'
sleep 10
if [$? -eq 0]; then
    echo "Tables populated successfully!!"
else
    echo "Failed to create tables"
    exit 1
fi

sleep 5
exit 0
