## Piemapping

## Dependencies

- go (1.8 used for dev)
- glide
- docker

## Setup

`glide install`

## Run

`make docker_tests` - build the cotainer and run all tests in a dev container

If you want to run the service locally you can `export` the variables in `docker/piemapping.dev.env` and `docker/mysql.env`, altering them to point to your local mysql instance, and then run:
```
$ make migrate_up

$ make docker build

$ make docker run
```
## Deployment

`make prod_build_container` - run tests in the dev container and the build the prod image

I built this microservice with testing in mind and utilised dependency injection to acheive maximum testability of components.

Inversion of control allows you to inject your artifacts in at the top most level of the app i.e. the 'main' package rather than importing them in each package. This affords you the ability to stub/mock certain components of the app at a top level.

Good examples of this are...
  - the service wide logger than can be turned on and off with the environment variable `PIEMAPPING_LOGGING`
  - the service wide logger can be configured to print a different date format with the environment variable `PIEMAPPING_LOG_FORMAT`
  - the mocks used in the various unit tests which are injected into their dependent components e.g. the `JourneyRepositoy` being injected into the http handlers.

## Method

When developing I...
  - ran the tests inside a container to get the environment as close to production as possible.
  - created json fixtures in a dedicated directory so as not to pollute test files with inline fixtures.
  - chose to lean on SQL as much as possible when retrieving the data because it's the right tool for the job.
  - implemented the repository pattern to abstract the http handlers away from data source specifics.
  - create an application wide logger for uniformity in code and output.
  - create a seperate function to easily seed the database.
  - configure the app through environment variables as this makes life easier when deploying to a container.
  - write tests in the order of their magnitude i.e. system tests then integration tests then unit tests and then develop each feature at a time.
  - adhere to the principles of the 'testing triangle' by writing minimal system tests, some integrations tests and lots of unit tests.
  - wrap http handlers so that generic logging doesn't pollute handler code.
  

  
