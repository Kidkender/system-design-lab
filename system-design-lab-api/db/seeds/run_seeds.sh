#!/bin/bash
# Run all seed files against the local dev database.
# Usage: bash db/seeds/run_seeds.sh
# Prerequisites: PostgreSQL running on port 5433, database system_db exists.

set -e

DB_URL="${DATABASE_URL:-postgres://root:root@localhost:5433/system_db}"

echo "==> Seeding Chat App scenario..."
psql "$DB_URL" -f db/seeds/chat_app.sql

echo "==> Seeding URL Shortener scenario..."
psql "$DB_URL" -f db/seeds/url_shortener.sql

echo "==> Done. Run 'GET /api/v1/scenarios' to verify."
