up:
	docker-compose -f docker-compose.yaml --env-file .env run app \
	go run main.go --remove-orphans
	
down:
	docker-compose down -v --remove-orphans && docker volume prune -f

apply-migrations:
	docker exec -it m migrate- database ${DB_URL} -path repository/migrations/ up

run-integration-tests:
	docker-compose -f docker-compose.yaml --env-file .env run app gotest -v ./...
