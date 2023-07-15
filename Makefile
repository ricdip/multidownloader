include config.mk

.DEFAULT_GOAL: help

.SILENT: help
.PHONY: help # show this help message
help:
	@grep -E '^.PHONY:.+#.+' Makefile | sed 's/.PHONY: //' | awk -F ' # ' '{printf "%-15s %s\n", $$1, $$2}'

.PHONY: install # install with go install
install: $(BIN)$(OBJ)
	$(GOINSTALL)

.PHONY: image # create docker image
image: build
	$(DOCKERBUILD) -f $(DOCKERFILE) -t $(IMAGENAME) .

.PHONY: rmi # remove docker image
rmi:
	$(DOCKERRMI) $(IMAGENAME)

.PHONY: build # build from sources
build: clean deps
	$(GOBUILD) -o $(BIN)$(OBJ)

.PHONY: clean # clean build file
clean:
	$(GOCLEAN) -i
	rm -rf $(BIN)

.PHONY: deps # install dependencies
deps:
	$(GOMODTIDY)

.PHONY: format # format go source files
format:
	$(GOFMT) ./...