.POSIX:

DESTDIR=dist
HUGO_VERSION=0.84.0

help: ## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

build: check | $(DESTDIR) ## Check version of Hugo, install if not present, and build website
	@hugo

$(DESTDIR):
	mkdir -p $(DESTDIR)

clean: ## Remove all files and directories in dist
	@rm -rf dist/*

post: ## Create a new post with POST_NAME and POST_TITLE provided
	@hugo new posts/$(POST_NAME).md
	@sed -i 's|title:.*|title: "$(POST_TITLE)"|' content/posts/$(POST_NAME).md

check: ## Check/analyze lint of the POSTS Markdown source file
	@markdownlint ./content/posts/*.md
	@markdown-link-check ./content/posts/*.md

validate: | $(DESTDIR) ## Check Holberton's W3C Validator (non-blocking quality indicator)
	@chmod +x ./w3c_validator.py
	@python3 w3c_validator.py ./dist/index.html || true

.PHONY: help build clean post check validate
