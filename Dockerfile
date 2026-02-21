FROM golang:1.25.7-alpine3.22 AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum /
RUN go mod download

COPY .env .
COPY . .
RUN go build -o /usr/local/bin/app ./...

ENTRYPOINT ["app"]