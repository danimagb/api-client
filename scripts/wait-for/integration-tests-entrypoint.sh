#!/bin/sh

# Abort on any error (including if wait-for-it fails).
set -e

# Wait for the backend to be up, if we know where it is.
 if [ -n "$API_URL" ]; then
  /app/scripts/wait-for.sh "$API_URL/$API_HEALTHCHECK"
 fi

# Run the main container command.
exec "$@"