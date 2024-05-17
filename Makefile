SUB_DIRS = protobuf migrations
PACKAGES	?= $(shell go list ./...)

all: $(SUB_DIRS)

$(SUB_DIRS):
	make -C $@

test:
	-@mkdir build
	go test $(PACKAGES) -v -cover -failfast

clean:
	-rm -rf build

.PHONY: $(SUB_DIRS)