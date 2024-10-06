#!/bin/sh
set -e

export MIGRATION_DIR=/migrations
# Run migrations
/bin/migrate up

# Start the server
/bin/server