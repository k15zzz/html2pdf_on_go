FROM golang:1.20

RUN apt-get update
RUN apt-get install -y wkhtmltopdf

WORKDIR /src

COPY .. /src

ENV PROJECT_DIR=/src \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    config=docker

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 80 80

ENTRYPOINT ["wkhtmltopdf"]
ENTRYPOINT  CompileDaemon --build="go build -o /src/bin/prod /src/cmd/app/main.go" --command="/src/bin/prod"