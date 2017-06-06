FROM golang:1.8
MAINTAINER Luc CHMIELOWSKI <luc.chmielowski@gmail.com>

# Envs
ENV GO15VENDOREXPERIMENT=1

EXPOSE 5000

RUN mkdir -p /go/src/github.com/iochti/auth-service
WORKDIR /go/src/github.com/iochti/auth-service
COPY . /go/src/github.com/iochti/auth-service

RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

RUN rm -rf /go/src/github.com/iochti/auth-service

CMD ["auth-service"]
