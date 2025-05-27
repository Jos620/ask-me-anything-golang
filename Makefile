generate:
	go generate ./...

up:
	bash ./scripts/start-api.sh

down:
	bash ./scripts/stop-db.sh