FROM golang:1.20-alpine AS build
# MAINTAINER Farhan Rivaldy <fariv.fariv12@gmail.com>

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN  ls /app

COPY . .


RUN go build -o main .

WORKDIR /app
CMD ["./main"]

EXPOSE 1323