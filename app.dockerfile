# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./
COPY core ./core/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /cakeauth

CMD ["/cakeauth"]