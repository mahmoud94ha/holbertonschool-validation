

build: ## to generate
	HUGO

post: ## create a new blog
	$(HUGO) new posts/$(POST_NAME).md
	echo '# $(POST_TITLE)' >> ./content/posts/$(POST_NAME).md

clean: ## to cleanup dist folder
	$(RM) dist/*
help: ## show commands functionality
	@grep "^[a-zA-Z\-\_0-9]*:.*##.*" $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
