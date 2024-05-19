SUB_DIRS = protobuf migrations
PACKAGES	?= $(shell go list ./...)

all: $(SUB_DIRS)

$(SUB_DIRS):
	make -C $@

test:
	go test $(PACKAGES) -v -cover -failfast

test_docker:
	-docker stop monify-test-postgres
	-docker rm monify-test-postgres
	docker run --name monify-test-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
	go test $(PACKAGES) -v -cover -failfast -tags docker

clean:

.PHONY: $(SUB_DIRS)

