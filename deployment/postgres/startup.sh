#!/bin/sh

set -e

# Define array of service names
SERVICES="game content leaderboard user"

# Wait for Postgres server to be ready
until nc -z -v -w30 postgres-service 5432
do
  echo "Waiting for Postgres server to start..."
  # wait for 2 seconds before checking again
  sleep 2
done

# Loop through services and create secrets and databases
for service in $SERVICES; do
    # Generate random username and password
    username=$(openssl rand -hex 6)
    password=$(openssl rand -hex 16)

    # Create Kubernetes Secret
    kubectl create secret generic ${service}-db-credentials \
    --from-literal=username=$username \
    --from-literal=password=$password

    # Create the user and database
    psql -h postgres-service -U postgres -c "CREATE USER $username WITH PASSWORD '$password';"
    psql -h postgres-service -U postgres -c "CREATE DATABASE ${service}_db OWNER $username;"
done

echo "Database initialization complete!"
