SHELL=/bin/bash
# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_VET=$(GO_CMD) vet

#測試
COVER_PROFILE=cover.out

#docker
DOCKER_CMD=docker
DOCKER_BUILD=$(DOCKER_CMD) build
DOCKER_PUSH=$(DOCKER_CMD) push
DOCKER_IMAGE_REGISTRY=registry.digiwincloud.com/ops/
DOCKER_IMAGE_NAME=drs

#打包
BINARY_NAME=drs
BINARY_UNIX=$(BINARY_NAME)_unix
VERSION:=$(shell cat VERSION)
PRE_BUILD_VERSION=$(shell cat BUILD)
SUB_VERSION:=$(shell echo ${PRE_BUILD_VERSION}+1 | bc)

#Git
GIT_CMD=git
GIT_BRANCH=$(GIT_CMD) branch
GIT_ADD=$(GIT_CMD) add
GIT_COMMIT=$(GIT_CMD) commit
GIT_PUSH=$(GIT_CMD) push
GIT_CURRENT_BRANCH=$(GIT_BRANCH) --show-current


#其他指令
ALL_PATH=./...

echo=echo
all: deps test build pack
deps:
	$(GO_GET) -u $(ALL_PATH)
test:
	$(GO_TEST) -v $(ALL_PATH) -cover
test_output:
	$(GO_TEST) -v $(ALL_PATH) -coverprofile=$(COVER_PROFILE)
build:
	$(GO_BUILD) -o $(BINARY_NAME)
run:
	./$(BINARY_NAME)
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
pack:
	tar -cvzf $(BINARY_NAME)-v$(VERSION).tar.gz $(BINARY_NAME) ./config ./i18n VERSION BUILD
docker_build:
	$(shell echo ${SUB_VERSION} > BUILD)
	$(GIT_ADD) BUILD
	$(GIT_COMMIT) -m "auto increase sub version"
	$(DOCKER_BUILD) -t $(DOCKER_IMAGE_REGISTRY)$(DOCKER_IMAGE_NAME):$(VERSION).$(SUB_VERSION) .
docker_push:
	$(DOCKER_PUSH) $(DOCKER_IMAGE_REGISTRY)$(DOCKER_IMAGE_NAME):$(VERSION).$(SUB_VERSION)
docker_ci: docker_build docker_push