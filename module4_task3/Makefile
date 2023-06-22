GO=go
TARGET=awesome-api
RM=rm -rf
MAKE=make
COVERAGE=coverage
HUGO=hugo
IMAGE_TAG=$(shell git tag | tail -n 1)
ifeq ($(strip $(IMAGE_TAG)),)
IMAGE_TAG=$(shell git rev-parse --abbrev-ref HEAD)
endif

build:	## Generate the source code of the application to a binary.
	docker run --rm -v $(PWD):/app -w /app awesome:build hugo --source . --destination dist
	docker run --rm -v $(PWD):/app -w /app awesome:build go build -o $(TARGET)

build-docker:	## Build docker image awesome:build
	docker build -t awesome:build -f build/Dockerfile .

docker-tests:   ## Test docker image awesome:build
	docker run --rm -v $(PWD):/app -v /var/run/docker.sock:/var/run/docker.sock -w /app awesome:build container-structure-test test --image awesome:build --config .build/cst.yml

run:	## Run the application in background by executing the binary.
	./$(TARGET) >./$(TARGET).log 2>&1 &

post:	## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME
	$(HUGO) new posts/$(POST_NAME).md
	sed -i '2s/.*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md

stop:	## Stop the application.
	@pkill $(TARGET)

clean:	## Stop the application. Delete the binary and the log file.
	@sudo $(RM) $(TARGET) $(TARGET).log $(COVERAGE)-units.out $(COVERAGE)-integrations.out dist/ *~

test:	## Test the application.
	$(MAKE) unit-tests
	$(MAKE) integration-tests
	$(MAKE) validate

unit-tests: ## make a coverage-units.out file
	touch coverage-units.out

integration-tests: ## make a coverage-integrations.out file
	touch coverage-integrations.out

lint: ## Lint Dockerfile
	docker run --rm -v $(PWD):/workdir -i hadolint/hadolint:v2.6.0 hadolint --config /workdir/cst.yml --no-fail /workdir/Dockerfile

check:	## Lint markdonw source, check dead links
	markdownlint ./content/posts/*.md
	find content -name \*.md -print0 | xargs -0 -n1 markdown-link-check

validate:	## Validate the file dist/index.html using the command line Holberton’s W3C Validator.
	docker run awesome:build sh -c "python3 ./w3c_validator.py ./dist/index.html || true"

package:	## Produce awesome-website.zip containing awesome-api and dist directory
	-make build
	docker build -t awesome:${IMAGE_TAG} .
	docker save -o awesome_image.tar awesome:${IMAGE_TAG}
	zip -FSr awesome-website.zip awesome-api dist awesome_image.tar

help:	## Show this help usage.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)
