FROM golang:1.26.4-alpine3.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o client .

CMD ["./client"]