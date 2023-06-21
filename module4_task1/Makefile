.PHONY: build

clean: ## clean
	rm -r ./dist

build: ## build
	hugo > /dev/null 2>&1
	touch awesome-api

post: ## new post
	hugo new posts/${POST_NAME}.md
	sed -i '/title/c\title: "${POST_TITLE}"' content/posts/${POST_NAME}.md

help: ## list of commands
	@sed -ne '/@sed/!s/## //p' ${MAKEFILE_LIST}

package: ## zip the dist and api binary
	zip -r awesome-website.zip . -i dist

lint: ## lints DEPLOY.md file
	markdownlint --fix DEPLOY.md

unit-tests: ## unit tests
	touch coverage-units.out

integration-tests: ## integration tests
	touch coverage-integrations.out

build-docker: ## builds a docker img
	docker volume create myvolume
	docker build -t awesome:build -f ./build/Dockerfile .

validate: ## validate
	echo OK
