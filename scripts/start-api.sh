#!/bin/bash

# Restart DB
bash scripts/stop-db.sh
bash scripts/start-db.sh

# Run API
go run ./cmd/ama/
