FROM ubuntu:16.04

RUN \
    apt-get update \
    && apt-get install -y \
        curl \
        git \
        make \
    && rm -rf /var/lib/apt/lists/*

RUN \
    curl -O https://godeb.s3.amazonaws.com/godeb-amd64.tar.gz \
    && tar xf godeb-amd64.tar.gz -C /usr/bin \
    && godeb install 1.8

RUN mkdir /go
ENV GOPATH /go
ENV PATH ${PATH}:${GOPATH}/bin

ENV WORKDIR ${GOPATH}/src/github.com/eggsbenjamin/piemapping
RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}
