install:
	@glide install

build:
	@go build -v -o ./bin/piemapping ./cmd

migrate_up:
	@./bin/piemapping migrate up

test:
	@make unit_tests && make integration_tests && make system_tests

unit_tests:
	@PIEMAPPING_LOGGING=false go test -v ./repository/... ./http_handlers/... -tags=unit

integration_tests:
	@PIEMAPPING_LOGGING=false go test -v ./repository/... -tags=integration

system_tests:
	@PIEMAPPING_LOGGING=false go test -v ./systemtest/...

docker_test:
	@docker-compose -f docker/docker-compose-tests.yml -p piemapping rm -v -f
	@docker-compose -f docker/docker-compose-tests.yml -p piemapping up --abort-on-container-exit

