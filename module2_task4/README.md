# Module 2

## Build an Application using Make

### Prerequisites
* Golang in v1.18
* golangci-lint

### Lifecycle
The life-cycle of this project is the following command:

 `make help`:
    - show all command description

 `build`:
    - Builds a new version of the website to folder `/dist/`

 `clean`:
    - Removes the contents the folder  `/dist/`

 `post`:
    - Creates a new post in the contents/post folder with POST_TITLE and POST_NAME set from the ENV variables.

 `check`:
	- Lint of the Markdown source files using command line AND analysis of the links with command line. If one test fails, the command failed.

 `validate`:
	- validate the file ./dist/index.html by using command line. But non-blocking quality indicator
