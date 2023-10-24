
VERSION :=1.0
run:
	@go run app/services/interest-cal-api/main.go

build:
	@go build -ldflags "-X main.build=local" ./app/services/interest-cal-api

docker:
	docker build \
		-f docker/dockerfile \
		-t interest-cal-api:${VERSION} \
		--build-arg BUILD_REF=${VERSION} \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
.PHONY: start
start: docker
	IMAGE_NAME=interest-cal-api:${VERSION} docker compose -f./docker/docker-compose.yml  up -d

.PHONY: stop
stop:
	IMAGE_NAME=interest-cal-api:${VERSION}  docker compose -f./docker/docker-compose.yml rm -v --force --stop
	
# docker image rm interest-cal-api:${VERSION}

.PHONY: run docker

