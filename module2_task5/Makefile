.POSIX:
DESTDIR=dist
HUGO_VERSION=0.84.0

help: ## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

hugo-build: ## Check version of Hugo, install if not available, and build website
	@make check && hugo -c $(DESTDIR)

go-build: ## Compile source code of the application into a binary named awesome.api
	@go mod init github.com/mahmoud94ha/awesome-api
	@go get github.com/gorilla/mux
	@go build -o awesome.api
	@make lint >/dev/null 2>&1 || exit -1

build: ## Execute Hugo and Golang builds
	@make hugo-build
	@make go-build

hugo-clean: ## Remove all files recursively (directories and content)
	@rm -rf $(DESTDIR)/*

go-clean: ## Stop the application and delete awesome-api and awesome.log
	@make stop
	@rm -rf awesome.api awesome.log coverage-units.out coverage-integrations.out

clean: ## Execute hugo-clean and go-clean
	@make hugo-clean
	@make go-clean

run: ## Run the application in the background and write logs to awesome-api.log
	@./awesome.api > ./awesome-api.log 2>&1 &

stop: ## Stop the application
	@kill $$(pgrep awesome.api) 2>/dev/null || true

post: ## Create a new post with POST_NAME and POST_TITLE
	@hugo new posts/$(POST_NAME).md
	@sed -i 's|title:.*|title: "$(POST_TITLE)"|' content/posts/$(POST_NAME).md

check: ## Check and analyze the lint of the Markdown source files
	@markdownlint ./content/posts/*.md
	@markdown-link-check ./content/posts/*.md

validate: | $(DESTDIR) ## Check Holberton's W3C Validator (non-blocking quality indicator)
	@chmod +x ./w3c_validator.py
	@python3 w3c_validator.py $(DESTDIR)/index.html || true

test: ## Run unit and integration tests for the API with code coverage
	@make unit-tests
	@make integration-tests

lint: ## Test lint on the files
	@golangci-lint run

unit-tests: ## Run Golang unit tests
	@go test -v -short -coverprofile=coverage-units.out ./...

integration-tests: ## Run Golang integration tests
	@go test -v -coverprofile=coverage-integrations.out ./...
