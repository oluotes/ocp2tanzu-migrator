# syntax=docker/dockerfile:1

##
## Build
##
##FROM golang:1.18-buster AS build
FROM golang:latest AS build

WORKDIR /app

COPY . ./
RUN go mod tidy

RUN go build -v

##
## Deploy
##
#FROM gcr.io/distroless/base-debian10
FROM golang:1.18-buster

WORKDIR /

COPY --from=build /app/db-sync-service /db-sync-service

EXPOSE 8080

USER 1001

ENTRYPOINT ["/db-sync-service"]
