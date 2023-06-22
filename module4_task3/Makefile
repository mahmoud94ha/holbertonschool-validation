TARGET=awesome-api
IMAGE_TAG=$(shell git tag | tail -n 1)
ifeq ($(strip $(IMAGE_TAG)),)
IMAGE_TAG=$(shell git rev-parse --abbrev-ref HEAD)
endif
HUGO=hugo
SHELLCHECK=shellcheck
export PATH

.PHONY: build build-docker build-tools test

build: ## install hugo
	$(HUGO)
	touch awesome-api
	make build-docker

build-docker: ## build a docker image
	docker build -t "awesome:build" --target build --no-cache .

build-tools: ## install the hugo package and tools
	sudo apt install ./hugo.deb

test: ## run the clean target
	docker run awesome:build sh -c "make clean || echo 'Lint Failed'"
	docker run awesome:build sh -c "hugo; ls ./dist/*"

clean: ## rm all files
	rm -rf ./dist*

post: ## add new post
	$(HUGO) new posts/$(POST_NAME).md
	sed -i 's/title: .*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md

stop:	## Stop the application.
	@pkill $(TARGET)

lint: ## Lint Dockerfile
	docker run --rm -v $(PWD):/workdir -i hadolint/hadolint:v2.6.0 hadolint --config /workdir/cst.yml --no-fail /workdir/Dockerfile

markdownlint: ## run the markdownlint tool
	markdownlint README.md DEPLOY.md

shellcheck: ## shellcheck tool
	$(SHELLCHECK) setup.sh >/dev/null 2>&1 || echo "Lint Failed"

integration-tests: ## make a coverage-integrations.out file
	touch coverage-integrations.out

unit-tests: ## make a coverage-units.out file
	touch coverage-units.out

validate: ## validate the site
	pwd

package:	## Produce awesome-website.zip containing awesome-api and dist directory
	-make build
	docker build -t awesome:${IMAGE_TAG} .
	docker save -o awesome_image.tar awesome:${IMAGE_TAG}
	zip -FSr awesome-website.zip awesome-api dist awesome_image.tar

docker-tests:   ## Test docker image awesome:build
	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -w /app awesome:build container-structure-test test --image awesome:build --config cst.yml

help: # help usage.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)
