up:
	docker-compose pull && docker-compose up --remove-orphans

apply-migrations:
	docker exec -it m migrate- database ${DB_URL} -path repository/migrations/ up

run-integration-tests:
	docker-compose --env-file .env run app gotest -v ./...
