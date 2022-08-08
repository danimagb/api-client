#!/bin/sh

# Abort on any error (including if wait-for-it fails).
set -e

# Wait for the backend to be up, if we know where it is.
 if [ -n "$HEALTH_CHECK_HOST" ]; then
  /app/scripts/wait-for-it.sh "$HEALTH_CHECK_HOST:${HEALTH_CHECK_PORT:-8080}"
 fi

# Run the main container command.
exec "$@"