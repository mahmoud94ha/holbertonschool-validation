help:## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

build:## Compile source code
	@go get github.com/gorilla/mux
	@make lint >/dev/null 2>&1 && go build || exit -1

run: ## Run the application in background
	@./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop
	@kill "$(shell pgrep awesome-api)"

clean:## clean
	@make stop
	@rm -rf awesome-api awesome-api.log

test: ## test to ensure expected beahaviour
	@curl http://localhost:9999
	@curl http://localhost:9999/health

lint: ## Test lint on the files
	@golangci-lint run

unit-tests: ## Run Golang uni tests
	@go test -v -short -coverprofile=coverage-units.out
