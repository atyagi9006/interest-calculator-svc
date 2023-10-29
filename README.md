# interest-calculator-svc
interest-calculator-svc is production grade service to calculate simple interest and store it on mongo DB

This service provides different feature like api extendability, reuseability, sepreration of concern

## Design Decision 
1. App layer mainly contains very specfic code to a different apis (i.e. in our case interest-cal-api) and which takes externl input. e.g 
2. Business layer coantains packages which used by the app layer to process the request related store and retrive data. It also some reusable components. 
3. There should be no circular imorts and imports are only done from down to top (i.e. any code in app packages could import any package in business layer.) But vicea-verse should not be done.    
4. This service could be build with different version. Need to sepecify in Makefile. Default value of VERSION  is 1.0

## Tech stack and framework used
1. GoFiber - for hosting service api handles
2. Mongodb - for storing data 
3. Docker - for containerization
4. counterfeiter- for generating mocks
5. testify - for test assertion 

## Design 
![Alt text](design.png)


# How to run Service

### **pre requisites**
1. Install docker and docker compose
2. install golang 

### **Steps to run**
To run service run this command in interest-calculator-svc.This first test the service and build the docker image will specified version and once done will spawn two containers (i.e.mongodb-container, interest-cal)
```
make start
```
To stop service run this command in interest-calculator-svc.This will stop and removed the running containers
```
make stop
```
To test service seperatly run this command. This will run all the table driven unit tests for service withour caching the last output and also checks the race condition. 
```
make test
```

To re-generate the mocks run this command.
```
make gen
```
To test the apis ```run make start``` and import the postmen collection in postamn or use ``` curl ``` to hit apis. All the api will hosted on ```localhost:3500```
# API Documentations
 HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| POST | /api/v1/si | To create simpleinterest record in mongodb  |
| GET | /api/v1/si/{id} | To get simpleinterest record from mongodb |
| DELETE | /api/v1/si/{id} | To delete simpleinterest record from mongodb |

## Project Tree
```
.
├── app
│   └── services
│       └── interest-cal-api
│           ├── handlers
│           │   ├── handler.go
│           │   └── v1
│           │       └── sigrp
│           │           └── sigrp.go
│           ├── main.go
│           └── start
│               └── start.go
├── bin
│   └── interest-cal-api
├── business
│   ├── core
│   │   ├── entities
│   │   │   └── simpleinterest.go
│   │   └── simpleInterest
│   │       ├── fakerepo
│   │       │   ├── fakeMongo
│   │       │   │   └── fake_repository.go
│   │       │   └── genfake.go
│   │       ├── repo
│   │       │   └── repo.go
│   │       ├── svc.go
│   │       └── svc_test.go
│   └── data
│       └── repoConn.go
├── docker
│   ├── docker-compose.yml
│   └── dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```