FROM golang:1.25.0-alpine
RUN go install github.com/air-verse/air@latest
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify





