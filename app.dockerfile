# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum main.go ./
COPY core ./core/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /cakeauth

CMD ["/cakeauth"]