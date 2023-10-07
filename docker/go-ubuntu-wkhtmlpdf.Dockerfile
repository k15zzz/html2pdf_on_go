FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get clean
RUN apt-get update

RUN apt-get install -y build-essential xorg libssl-dev libxrender-dev wget gdebi
RUN apt-get install -y wkhtmltopdf

RUN apt-get install -y wget git gcc

RUN wget -P /tmp "https://dl.google.com/go/go1.20.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go1.20.linux-amd64.tar.gz"
RUN rm "/tmp/go1.20.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

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
