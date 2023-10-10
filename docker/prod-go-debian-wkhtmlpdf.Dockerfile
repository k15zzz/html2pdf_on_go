FROM golang:1.20 AS build

WORKDIR /application

COPY . /application

ENV config=docker

RUN go mod download && \
    go mod verify && \
    go build -o /src/bin/app ./cmd/app/main.go

FROM debian:stable

RUN apt-get update && \
    apt-get install -y wkhtmltopdf && \
    apt-get clean

COPY --from=build /src/bin/app /source/service/app
COPY --from=build /application/configs/docker.yml /source/service/configs/docker.yml

CMD cd /source/service && ./app