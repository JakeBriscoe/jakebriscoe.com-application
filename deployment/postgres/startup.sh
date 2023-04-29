#!/bin/sh

set -e

# Wait for Postgres server to be ready
until pg_isready -h postgres-service -p 5432 -q
do
  echo "Waiting for Postgres server to start..."
  # wait for 2 seconds before checking again
  sleep 2
done

# Define array of service names
SERVICES="game content leaderboard user"

# Loop through services and create secrets and databases
for service in $SERVICES; do
    echo "Creating secrets and database for ${service}"

    # Generate random username and password
    username=$(openssl rand -hex 6)
    password=$(openssl rand -hex 16)

    # Create Kubernetes Secret
    kubectl create secret generic ${service}-db-credentials \
    --from-literal=username=$username \
    --from-literal=password=$password

    echo "Connecting to server with username ${POSTGRES_USER} and password ${PGPASSWORD} and db ${POSTGRES_DB}"

    # Create the user and database
    psql -h postgres-service -U $POSTGRES_USER -d $POSTGRES_DB -c "CREATE USER \"$username\" WITH PASSWORD '$password';"
    psql -h postgres-service -U $POSTGRES_USER -d $POSTGRES_DB -c "CREATE DATABASE ${service}_db OWNER '$username';"

    echo "Database created for ${service}"
done

echo "Database initialization complete!"

exit 0