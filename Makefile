up:
	docker-compose -f docker-compose.yaml --env-file .env up \
	--remove-orphans --build
	
down:
	docker-compose down -v --remove-orphans && docker volume prune -f

apply-migrations:
	docker exec -it m migrate- database ${DB_URL} -path repository/migrations/ up

run-integration-tests:
		docker volume prune -f && \
		docker-compose -f docker-compose.test.yaml build && \
		docker-compose -f docker-compose.test.yaml --env-file .env \
		run test_app gotest -run 'Integration' -v -p=1 ./...
