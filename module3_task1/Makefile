build: ## generate the website
		hugo

clean: ## cleanup dist directory
		rm -fr dist/*

post: ## create a new blog
		hugo new posts/$(POST_NAME).md
		sed -e 's/title: .*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md > aux.tmp && mv aux.tmp content/posts/$(POST_NAME).md

help: ## show commands functionality
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

greet: ## greets
	echo "Hello my friend"

