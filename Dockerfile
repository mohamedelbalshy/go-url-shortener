FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY docs/ docs/

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
