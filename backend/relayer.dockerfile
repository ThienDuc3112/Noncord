FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@v1.62

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY air_config/relayer.air.toml .air.toml

RUN mkdir -p /app/tmp/air && chmod -R 0777 /app/tmp/air
