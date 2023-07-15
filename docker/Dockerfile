FROM golang:alpine3.13

LABEL maintainer="Riccardo Armando Di Prinzio (ricdip)"

COPY . /go/src/multidownloader

RUN apk update && \
    apk add make && \
    cd /go/src/multidownloader && \
    make build && \
    make install

WORKDIR /downloads

ENTRYPOINT ["multidownloader"]
CMD ["--help"]
