.PHONY: staticd testd test up down


up:
	docker-compose --file docker-compose-test.yml down --remove-orphans 2>/dev/null 1>&2
	docker-compose --file docker-compose-test.yml up --build -d

down:
	docker-compose --file docker-compose-test.yml down --remove-orphans 2>/dev/null 1>&2

test:
	@ echo "Running tests for party-service"
	docker-compose --file docker-compose-test.yml exec party-service go test ./...

testd:
	make up
	make test
	make down

build:
    @ echo "Todo: command should run all static analysis, test, package and publish on success"
