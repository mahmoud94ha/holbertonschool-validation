# Module 2

## Build an Application using Make

### Prerequisites
* Golang in v1.18
* golangci-lint

### Lifecycle
The life-cycle of this project is the following command:

* ```build```  compile the source code of the application to a binary named ```awesome-api``` (the name ```awesome-api``` comes from the command ```go mod init github.com/<your github handle>/awesome-api```) with the command go build. The first build may takes some times. Build run only if lint is not failed.

* ```run```: Run the application in background by executing the binary ```awesome-api```, and write logs into a file named ```awesome-api.log```

* ```stop```: Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application.

* ```clean```: Stop the application. Delete the binary ```awesome-api``` and the log file ```awesome-api.log```, ```coverage-untis.out``` and ```coverage-integrations.out```

* ```test```: Test to ensure that it behaves as expected.

* ```lint```: Test lint in the files

* ```unit-tests```: Run files with the _test.go suffix

* ```integration-tests```: Run files with the _integration_test.go suffix

* ```help```: print a list of all the goals
