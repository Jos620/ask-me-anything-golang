start-db:
	bash ./scripts/start-db.sh

generate-pgdb:
	sqlc generate -f ./configs/sqlc.yaml
