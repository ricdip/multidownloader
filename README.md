# Multidownloader

This program downloads the links passed as arguments.
Multidownloader is written in Go and can be containerized using Docker.

## Build program

```bash
user@host:~$ make all
```

## Create Docker image

```bash
user@host:~$ make image
```

## Delete Docker image

```bash
user@host:~$ make rmi
```

## Use Multidownloader from host

```bash
user@host:~$ bin/multidownloader <link_1> <link_2>
```

## Use Multidownloader from Docker container

```bash
user@host:~$ docker run -v $(pwd):/downloads -it ricdip/multidownloader:latest <link_1> <link_2>
```
