# Multidownloader

This program downloads the links passed as arguments.
Multidownloader is written in Go and can be containerized using Docker.

## Build program

```bash
make all
```

## Create Docker image

```bash
make image
```

## Delete Docker image

```bash
make rmi
```

## Use Multidownloader from host

```bash
bin/multidownloader <link_1> <link_2>
```

## Use Multidownloader from Docker container

```bash
docker run -v $(pwd):/downloads -it ricdip/multidownloader:latest <link_1> <link_2>
```
