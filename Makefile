.PHONY: staticd testd test

test:
	@ echo "Running tests for events-service"
	docker-compose --file docker-compose-test.yml exec events-service go test
	# Todo: drop test database

testd:
	docker-compose --file docker-compose-test.yml down --remove-orphans 2>/dev/null 1>&2
	docker-compose --file docker-compose-test.yml up --build -d
	make test
	docker-compose --file docker-compose-test.yml down --remove-orphans 2>/dev/null 1>&2

