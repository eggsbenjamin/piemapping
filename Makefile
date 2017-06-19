install:
	@glide install

build:
	@go build -v -o ./bin/piemapping ./cmd

run:
	@./bin/piemapping run

migrate_up:
	@./bin/piemapping migrate up

test:
	@PIEMAPPING_LOGGING=false make unit_tests && make integration_tests && make system_tests

unit_tests:
	@PIEMAPPING_LOGGING=false go test -v ./repository/... ./http_handlers/... -tags=unit

integration_tests:
	@PIEMAPPING_LOGGING=false go test -v ./repository/... -tags=integration

system_tests:
	@go test -v ./systemtest/...

docker_test:
	@docker-compose -f docker/docker-compose-tests.yml -p piemapping rm -v -f
	@docker-compose -f docker/docker-compose-tests.yml -p piemapping up --abort-on-container-exit

prod_build_binary:
	@CGO_ENABLED=0 GOOS=linux go build -v -o ./bin/piemapping ./cmd
	@chmod -R 0777 bin

prod_build_container:
	@make docker_test && make prod_build_binary && docker build -t piemapping .
	
