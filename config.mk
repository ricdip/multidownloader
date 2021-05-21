GO = go
GOBUILD = $(GO) build
GOMOD = $(GO) mod
GOINSTALL = $(GO) install
GOMODTIDY = $(GOMOD) tidy
DOCKER = docker
DOCKERBUILD = $(DOCKER) build
DOCKERRMI = $(DOCKER) rmi

OBJ = multidownloader
BIN = bin/
IMAGENAME = ricdip/multidownloader:latest
