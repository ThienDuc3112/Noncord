#! /bin/sh

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://noncord:password@localhost:6543/noncord?sslmode=disable"
export GOOSE_MIGRATION_DIR=internal/infra/db/sql/migration

goose "$@"
