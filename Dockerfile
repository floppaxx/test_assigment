FROM golang:1.20.5

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build -o app

CMD ["./app"]
