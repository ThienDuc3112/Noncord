#! /bin/sh

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://postgres:password@localhost:5556/chat?sslmode=disable"
export GOOSE_MIGRATION_DIR=internal/infra/db/sql/migration

goose "$@"
