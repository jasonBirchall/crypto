FROM golang:1.15-alpine AS build

WORKDIR /app
ENV CGO_ENABLED=0
ARG VERSION

COPY . .
RUN go build -ldflags=-s -ldflags=-w -ldflags=-X=github.com/jasonbirchall/crypto/crypto.version=$VERSION -o crypto .
ENV CGO_ENABLED=0
