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

    # Check if secret already exists
    if kubectl get secret ${service}-db-credentials >/dev/null 2>&1; then
        echo "Secret ${service}-db-credentials already exists"
        continue
    fi

    # Generate random username and password
    username=$(openssl rand -hex 6)
    password=$(openssl rand -hex 16)

    # Create the user and database
    if psql -h postgres-service -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "CREATE USER \"$username\" WITH PASSWORD '$password';"; then
      echo "User created for ${service}"
    else
      echo "Failed to create user for ${service}"
      continue
    fi

    if psql -h postgres-service -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "CREATE DATABASE ${service}_db OWNER '$username';"; then
      echo "Database created for ${service}"
    else
      echo "Failed to create database for ${service}"
      continue
    fi

        # Create Kubernetes Secret
    if kubectl create secret generic ${service}-db-credentials --from-literal=username="$username" --from-literal=password="$password"; then
      echo "Secret created for ${service}"
    else
      echo "Failed to create secret for ${service}"
      continue
    fi
done

echo "Database initialization complete!"

exit 0
