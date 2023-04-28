#!/bin/bash
set -e

# Define array of service names
services=("game" "content" "leaderboard" "user")

# Loop through services and create secrets and databases
for service in "${services[@]}"; do
    # Generate random username and password
    username=$(openssl rand -hex 6)
    password=$(openssl rand -hex 16)

    # Create Kubernetes Secret
    kubectl create secret generic ${service}-db-credentials \
    --from-literal=username=$username \
    --from-literal=password=$password

    # Can also write these tp docker-entrypoint-initdb.d/01-create-dbs.sh 
    # which will auto run on init
    # Create the user and database
    psql -U postgres -c "CREATE USER $username WITH PASSWORD '$password';"
    psql -U postgres -c "CREATE DATABASE ${service}_db OWNER $username;"
done
