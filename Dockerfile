FROM golang:1.16-buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD go run cmd/goth/main.go
