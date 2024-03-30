FROM golang:1.22-alpine

WORKDIR /app

COPY . .

run go mod tidy
run go build -o main main.go

# EXPOSE 3000

# COPY main.go /

# CMD["./main"]

CMD ["./main"]
