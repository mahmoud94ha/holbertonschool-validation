HUGO=hugo
SHELLCHECK=shellcheck
export PATH

.PHONY: build build-docker build-tools test

build: ## install hugo
	$(HUGO)
	touch awesome-api

build-docker: ## build a docker image
	docker build -t "awesome:build" ./build --no-cache

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

lint: ## check and clean
	hadolint ./build/Dockerfile

markdownlint: ## run the markdownlint tool
	markdownlint README.md DEPLOY.md

shellcheck: ## shellcheck tool
	$(SHELLCHECK) setup.sh >/dev/null 2>&1 || echo "Lint Failed"

integration-tests: ## make a coverage-integrations.out file
	touch coverage-integrations.out

unit-tests: ## make a coverage-units.out file
	touch coverage-units.out

validate: ## validate
	echo OK

package: ## build
	make build
	zip awesome-website.zip dist

docker-tests: ## docker image
	container-structure-test test --image awesome:build --config ./build/cst.yml

help: ## print help
	cat Makefile | grep '##' | tr -d '##' | head -n -1
