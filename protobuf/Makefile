MODULES = $(wildcard *.proto)

OUTPUT = $(MODULES:%.proto=%.pb.go)
OUTPUT += $(MODULES:%.proto=%_grpc.pb.go)

# Default target
all: $(OUTPUT)

%.pb.go: %.proto
	protoc --go_out=. --go_opt=paths=source_relative $<
%_grpc.pb.go: %.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative $<

force_rebuild:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto