MODULES = $(wildcard *.proto)
SWAGGER_MODULES = group.proto user.proto auth.proto group_bill.proto
O_DIR = gen/go
SWAGGER_DIR = gen
OUTPUT = $(MODULES:%.proto=$(O_DIR)/%.pb.go)
OUTPUT += $(MODULES:%.proto=$(O_DIR)/%_grpc.pb.go)
OUTPUT += $(MODULES:%.proto=$(O_DIR)/%.pb.gw.go)

# Default target
all: $(OUTPUT) $(SWAGGER_DIR)/monify.swagger.json
	mkdir -p gen/go



$(SWAGGER_DIR)/monify.swagger.json:
	-mkdir $(SWAGGER_DIR)
	protoc -I . --openapiv2_out $(SWAGGER_DIR) --openapiv2_opt=preserve_rpc_order=true \
  		--openapiv2_opt=allow_merge=true \
  		--openapiv2_opt=disable_default_errors=true \
  		--openapiv2_opt=merge_file_name=monify $(SWAGGER_MODULES)

$(O_DIR)/%.pb.go: %.proto $(O_DIR)
	protoc --go_out=./gen/go --go_opt=paths=source_relative $<

$(O_DIR)/%_grpc.pb.go: %.proto $(O_DIR)
	protoc --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative $<

$(O_DIR)/%.pb.gw.go: %.proto $(O_DIR)
	protoc -I . --grpc-gateway_out ./gen/go  --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true $<

$(O_DIR):
	@mkdir -p $(O_DIR)

force_rebuild:
	protoc --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative *.proto
