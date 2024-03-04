FROM golang:1.22-alpine

WORKDIR /app

COPY . .

run go mod tidy
run go build -o bin/main main.go


CMD["./bin/main"]

