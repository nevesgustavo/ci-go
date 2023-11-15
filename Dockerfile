FROM golang:1.21

WORKDIR /app

RUN go mod init teste

COPY .github .

RUN go build -o math

CMD ["./math"]