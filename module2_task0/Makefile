help:## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

build:## Compile source code
	@go get github.com/gorilla/mux
	@go build

run: ## Run the application in background
	@./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop
	@kill "$(shell pgrep awesome-api)"

clean:## clean
	@make stop
	@rm -rf awesome-api

test: ## test to ensure expected beahaviour
	@curl http://localhost:9999
	@curl http://localhost:9999/health
