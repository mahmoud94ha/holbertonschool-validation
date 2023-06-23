# Module 2

## Build an Application using Make

### Prerequisites
* Golang in v1.18
* golangci-lint

### Lifecycle
The life-cycle of this project is the following command:

To execute the Makefile use the following syntax:
 ```make <command>```

 command are availaible :
* `help`:
    - show all command description

* `hugo-build`:
    - Builds a new version of the website to folder `/dist/`

* `go-build`:
	- compile the source code of the application to a binary named ```awesome-api``` (the name ```awesome-api``` comes from the command ```go mod init github.com/<your github handle>/awesome-api```) with the command go build. The first build may takes some times. Build run only if lint is not failed.

* `build`:
	- execute `hugo-build` AND `go-build`

* `hugo-clean`:
    - Removes the contents the folder  `/dist/`

* `go-clean`:
	- Stop the application. Delete the binary ```awesome-api``` and the log file ```awesome-api.log```

* `clean`:
	- execute `hugo-clean` AND `go-clean`

* `post`:
    - Creates a new post in the contents/post folder with POST_TITLE and POST_NAME set from the ENV variables.

* `check`:
	- Lint of the Markdown source files using command line AND analysis of the links with command line. If one test fails, the command failed.

* `validate`:
	- validate the file ./dist/index.html by using command line. But non-blocking quality indicator

* `test`: Test to ensure that it behaves as expected.

* `lint`: Test lint in the files

* `unit-tests`: Run files with the _test.go suffix

* `integration-tests`: Run Golang integration tests

* `run`: Run the application in background by executing the binary ```awesome-api```, and write logs into a file named ```awesome-api.log```

* `stop`: Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application.
