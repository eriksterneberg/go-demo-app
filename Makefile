.PHONY: staticd testd test

test:
	@ echo "Running tests for events-service"
	@ docker-compose run events-service go test ./events-service/src

testd:
	@ docker-compose down --remove-orphans 2>/dev/null 1>&2
	@ docker-compose up --build -d web 2>/dev/null 1>&2
	@ make test
	@ docker-compose down --remove-orphans 2>/dev/null 1>&2

