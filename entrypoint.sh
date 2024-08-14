#!/bin/sh

# Migrate database
migrate -database postgres://postgres:password@postgres:5432/db_social_service -path external/database/migrations up

# Running app
./go-social-service
