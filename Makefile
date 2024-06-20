SUB_DIRS = protobuf
PACKAGES	?= $(shell go list ./...)

all: $(SUB_DIRS)

$(SUB_DIRS):
	make -C $@

test:
	-mkdir build
	go test $(PACKAGES) -v -cover -failfast

test_docker:
	-docker stop monify-test-postgres
	-docker rm monify-test-postgres
	docker run --name monify-test-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
	go test $(PACKAGES) -v -cover -failfast -tags docker

clean:
	-rm -rf build


docker_push_proxy: docker_build_proxy
	docker push registry.nccupass.com/monify_restful_proxy
docker_push_monify: docker_build_monify
	docker push registry.nccupass.com/monify
docker_push_media: docker_build_media
	docker push registry.nccupass.com/media_service
docker_push: docker_push_proxy docker_push_monify docker_push_media
docker_build: docker_build_proxy docker_build_monify
docker_build_proxy:
	 docker build -f Dockerfile.proxy -t registry.nccupass.com/monify_restful_proxy .
docker_build_monify:
	 docker build -f Dockerfile.api -t registry.nccupass.com/monify .
docker_build_media:
	 docker build -f Dockerfile.media -t registry.nccupass.com/media_service .

.PHONY: $(SUB_DIRS)