install:
	@glide install

unit_tests:
	@go test -v ./repository/... -tags=unit

integration_tests:
	@go test -v ./repository/... -tags=integration

system_tests:
	@go test -v ./systemtest/...
