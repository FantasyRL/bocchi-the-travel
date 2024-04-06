FROM golang:1.21 AS builder
LABEL authors="fanr"

ENV TZ Asia/Shanghai
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go mod tidy
#RUN make build-all