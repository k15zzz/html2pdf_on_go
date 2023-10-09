# image: k15zzz/go-debian-wkhtmlpdf:1.20
FROM golang:1.20

RUN apt-get update
RUN apt-get install -y wkhtmltopdf

COPY . .

ENV config=docker

EXPOSE 80 80

ENTRYPOINT ["wkhtmltopdf"]
ENTRYPOINT  go run ./cmd/app/main.go