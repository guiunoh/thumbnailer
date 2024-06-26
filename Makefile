.PHONY:

pkg = thumbnailer

# go test -coverprofile=coverage.out -count=1 ./...
# go tool cover -html=coverage.out -o cover.html
coverage:
	go test -cover -count=1 ./...

test:
	go test -count=1 ./...

run:
	go run ./cmd/$(pkg)

dockerize:
	docker build \
	--tag $(pkg) \
	--build-arg package=./cmd/$(pkg) \
	--build-arg config=config.yaml \
	.

run-docker:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:8081 \
	$(pkg)

run-docker-with-mount:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	--mount type=bind,source=$(PWD)/config/config.dev.yaml,target=/home/deploy/config/config.yaml,readonly \
	$(pkg)

run-docker-with-config:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	--env APP_CONFIG_FILE=config.dev.yaml \
	$(pkg)
