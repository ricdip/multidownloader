include config.mk

.DEFAULT_GOAL: help

.SILENT: help
.PHONY: help # show this help message
help:
	@grep -E '^.PHONY:.+#.+' Makefile | sed 's/.PHONY: //' | awk -F ' # ' '{printf "%-15s %s\n", $$1, $$2}'

.PHONY: all # clean, install deps, build
all: clean deps build

.PHONY: install # install with go install
install: $(BIN)$(OBJ)
	$(GOINSTALL)

.PHONY: image # create docker image
image: build
	$(DOCKERBUILD) -t $(IMAGENAME) .

.PHONY: rmi # remove docker image
rmi:
	$(DOCKERRMI) $(IMAGENAME)

.PHONY: build # build from sources
build:
	$(GOBUILD) -o $(BIN)$(OBJ)

.PHONY: clean # clean build file
clean:
	rm -f $(BIN)$(OBJ)

.PHONY: deps # install dependencies
deps:
	$(GOMODTIDY)
