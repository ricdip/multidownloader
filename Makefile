.POSIX:

include config.mk

.DEFAULT_GOAL: help


.SILENT: help
help:
	printf "%s \\t\\t %s\n" "help" "show this help message"
	printf "%s \\t\\t %s\n" "all" "clean, install deps, build"
	printf "%s \\t %s\n" "install" "install with go install"
	printf "%s \\t\\t %s\n" "image" "create docker image"
	printf "%s \\t\\t %s\n" "rmi" "remove docker image"
	printf "%s \\t\\t %s\n" "build" "build from sources"
	printf "%s \\t\\t %s\n" "clean" "clean build file"
	printf "%s \\t\\t %s\n" "deps" "install dependencies"

all: clean deps build

install: $(BIN)$(OBJ)
	$(GOINSTALL)

image: build
	$(DOCKERBUILD) -t $(IMAGENAME) .

rmi:
	$(DOCKERRMI) $(IMAGENAME)

build:
	$(GOBUILD) -o $(BIN)$(OBJ)

clean:
	rm -f $(BIN)$(OBJ)

deps:
	$(GOMODTIDY)

.PHONY: help all install image rmi build clean deps
