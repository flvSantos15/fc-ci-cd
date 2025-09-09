FROM golang:latest

WORKDIR /app

RUN go mod init teste

COPY . .

RUN go build -o math

CMD ["./math"]
