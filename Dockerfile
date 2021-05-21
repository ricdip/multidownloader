FROM golang:alpine3.13

LABEL maintainer="riccardodpr@gmail.com"

COPY . /go/src/multidownloader

RUN apk update && \
    apk add make && \
    cd /go/src/multidownloader && \
    make all && \
    make install

WORKDIR /downloads

ENTRYPOINT ["downloader"]
CMD ["--help"]
