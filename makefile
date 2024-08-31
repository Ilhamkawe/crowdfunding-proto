GO_MODULE := github.com/Ilhamkawe/crowdfunding-proto

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
	mkdir protogen\go
else
	rm -fR ./protogen
	mkdir -p ./protogen/go
endif

.PHONY: protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go_grpc_opt=module=${GO_MODULE} --go_grpc_out=. \
	./proto/Auth/*.proto

.PHONY: build
build: clean protoc-go

.PHONY: pipeline-init
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: pipeline-build
pipeline-build : pipeline-init build