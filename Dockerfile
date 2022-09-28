# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.19.1-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD [ "/app/main" ]
