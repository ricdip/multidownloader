# Multidownloader

This program downloads the links passed as arguments.
Multidownloader is written in Go and can be containerized using Docker.

## Makefile help message

```
help            show this help message
install         install with go install
image           create docker image
rmi             remove docker image
build           build from sources
clean           clean build file
deps            install dependencies
format          format go source files
```

## Multidownloader help message

```
Downloads the links passed as arguments

Usage:
  multidownloader [flags]

Flags:
  -c, --color              enable colored logging
  -d, --debug              enable debug logging
  -h, --help               help for multidownloader
  -l, --link stringArray   each link to download (-l <link_1> -l <link_2> ...)
  -q, --quiet              enable quiet logging (implies --yes)
  -v, --version            version for multidownloader
  -y, --yes                automatic yes to prompt
```

## Build

```bash
user@host:~$ make build
```

## Install

```bash
user@host:~$ make install
```

## Create Docker image

```bash
user@host:~$ make image
```

## Delete Docker image

```bash
user@host:~$ make rmi
```

## Run from host (after install)

```bash
user@host:~$ multidownloader -l <link_1> -l <link_2>
```

## Run from Docker container

```bash
user@host:~$ docker run --rm -v $(pwd):/downloads -it ricdip/multidownloader:latest -l <link_1> -l <link_2>
```
