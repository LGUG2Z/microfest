EXECUTABLES = git go swagger find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: clean vendor generate fmt install

generate:
	swagger generate server -f swagger.yml -P models.Principal

vendor:
	go mod vendor

install:
	go install ${LDFLAGS} ./cmd/microfest-server

fmt:
	gofmt -s -w cmd handlers models restapi
	goimports -w cmd handlers models restapi

release:
	goreleaser --rm-dist

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete
	rm -rf dist

.PHONY: check clean install build_all all
