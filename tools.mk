# Directories.
ROOT_DIR        :=$(shell git rev-parse --show-toplevel)
TOOLS_DIR       = $(abspath $(ROOT_DIR)/hack/tools)
SCRIPTS_DIR     = $(abspath $(ROOT_DIR)/scripts)
TOOLS_BIN_DIR   = $(TOOLS_DIR)/bin

GO_INSTALL      = $(abspath $(ROOT_DIR)/scripts/go_install.sh)

# Binaries.
GOFMT               = gofmt
GO                  = go
GOLANGCI_LINT       = $(TOOLS_BIN_DIR)/golangci-lint
CONTROLLER_GEN      = $(TOOLS_BIN_DIR)/controller-gen
CONVERSION_GEN      = $(TOOLS_BIN_DIR)/conversion-gen
KUBECTL             = $(TOOLS_BIN_DIR)/kubectl
KUBE_APISERVER      = $(TOOLS_BIN_DIR)/kube-apiserver
ETCD                = $(TOOLS_BIN_DIR)/etcd
KUBEBUILDER         = $(TOOLS_BIN_DIR)/kubebuilder
KIND                = $(TOOLS_BIN_DIR)/kind
KUSTOMIZE           = $(TOOLS_BIN_DIR)/kustomize
GOLINT              = $(TOOLS_BIN_DIR)/golint
GOX                 = $(TOOLS_BIN_DIR)/gox
GCOV2LCOV           = $(TOOLS_BIN_DIR)/gcov2lcov

# Common paths
REGISTRY ?= localhost:5000/fake

$(KIND): ## Install kind tool
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) sigs.k8s.io/kind@v0.8.1

$(KUSTOMIZE): ## Install kustomize
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) sigs.k8s.io/kustomize/kustomize/v3@v3.5.4

$(CONTROLLER_GEN): ## Build controller-gen from tools folder.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.0

$(CONVERSION_GEN): ## Build conversion-gen from tools folder.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/conversion-gen@v0.18.2

$(GOLANGCI_LINT): ## Build golangci-lint from tools folder.
	echo $(ROOT_DIR)
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint@v1.29.0

.PHONY: header-check
header-check: ## Runs header checks on all files to verify boilerplate
	$(SCRIPTS_DIR)/verify_boilerplate.sh

# TODO: are these used?
$(GOLINT): ## Build golint
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) golang.org/x/lint/golint

$(GOX): ## Build gox
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL)  github.com/mitchellh/gox@v1.0.1

$(GCOV2LCOV): ## Build gcov2lcov
	# NOTE: if you update the version here, also update ./github/workflows/test.yml
	#       for the Windows part of the "test-generator" job.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/jandelgado/gcov2lcov@v1.0.2

.PHONY: install-tools
install-tools: $(KIND) $(KUSTOMIZE) $(CONTROLLER_GEN) $(CONVERSION_GEN) $(GOLANGCI_LINT) $(GOLINT) $(GOX) $(GCOV2LCOV)

$(KUBECTL) $(KUBE_APISERVER) $(ETCD) $(KUBEBUILDER): ## Install test asset kubectl, kube-apiserver, etcd
	. $(SCRIPTS_DIR)/fetch_ext_bins.sh && fetch_tools
