
GENERATED_PATH=generated
GENERATED_API_PATH=$(GENERATED_PATH)/proto

all: install generate-proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-proto:
	if [ -d "$(GENERATED_API_PATH)" ]; then \
		rm -rf $(GENERATED_API_PATH); \
	fi
	mkdir $(GENERATED_API_PATH)
	protoc \
		--go_out=$(GENERATED_PATH) --go_opt=paths=source_relative \
    	--go-grpc_out=$(GENERATED_PATH) --go-grpc_opt=paths=source_relative \
    	proto/api.proto