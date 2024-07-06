.PHONY: proto
proto: protoc-gen-go protoc-gen-go-grpc protoc ## Generate proto files.
	PATH=$(PATH):$(LOCALBIN) protoc \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go_out=. --go-grpc_out=. external_api/proto/v1/*.proto

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
PROTOC_GEN_GO ?= $(LOCALBIN)/protoc-gen-go-$(PROTOC_GEN_GO_VERSION)
PROTOC_GEN_GO_GRPC ?= $(LOCALBIN)/protoc-gen-go-grpc-$(PROTOC_GEN_GO_GRPC_VERSION)

PROTOC_GEN_GO_VERSION ?= v1.28
PROTOC_GEN_GO_GRPC_VERSION ?= v1.2

.PHONY: protoc-gen-go
protoc-gen-go: $(PROTOC_GEN_GO) ## Download protoc-gen-go locally if necessary.
$(PROTOC_GEN_GO): $(LOCALBIN)
	$(call go-install-tool,$(PROTOC_GEN_GO),google.golang.org/protobuf/cmd/protoc-gen-go,${PROTOC_GEN_GO_VERSION})

.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC) ## Download protoc-gen-go-grpc locally if necessary.
$(PROTOC_GEN_GO_GRPC): $(LOCALBIN)
	$(call go-install-tool,$(PROTOC_GEN_GO_GRPC),google.golang.org/grpc/cmd/protoc-gen-go-grpc,${PROTOC_GEN_GO_GRPC_VERSION})

.PHONY: protoc
protoc:
	sudo apt-get install -y protobuf-compiler

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef