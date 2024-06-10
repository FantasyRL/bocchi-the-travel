FROM golang:1.22 AS builder
LABEL authors="fanr"

ENV TZ Asia/Shanghai
RUN go env -w GO111MODULE=on \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GOOS=linux \
  && go env -w GOARCH=amd64
SHELL ["/bin/bash", "-c"]

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go mod tidy
RUN make build-all
