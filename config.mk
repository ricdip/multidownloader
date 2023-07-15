GO = go
GOBUILD = $(GO) build
GOMOD = $(GO) mod
GOFMT = $(GO) fmt
GOINSTALL = $(GO) install
GOMODTIDY = $(GOMOD) tidy
DOCKER = docker
DOCKERBUILD = $(DOCKER) build
DOCKERRMI = $(DOCKER) rmi

OBJ = multidownloader
BIN = bin/
DOCKERFILE = docker/Dockerfile
IMAGENAME = ricdip/multidownloader:latest
