
VERSION :=1.0
.PHONY: run
run:
	@go run app/services/interest-cal-api/main.go

.PHONY: build
build:
	@go build -o bin/ -ldflags "-X main.build=local" ./app/services/interest-cal-api 

.PHONY:docker
docker:
	docker build \
		-f docker/dockerfile \
		-t interest-cal-api:${VERSION} \
		--build-arg BUILD_REF=${VERSION} \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
		
.PHONY: start
start: test docker
	IMAGE_NAME=interest-cal-api:${VERSION} docker compose -f./docker/docker-compose.yml  up -d

.PHONY: stop
stop:
	IMAGE_NAME=interest-cal-api:${VERSION}  docker compose -f./docker/docker-compose.yml rm -v --force --stop
# docker image rm interest-cal-api:${VERSION}

.PHONY: gen
gen:
	go get -d github.com/maxbrunsfeld/counterfeiter/v6
	go generate ./...

.PHONY: test
test:
	go mod tidy
	go test -v ./... -count=1 --race