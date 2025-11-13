FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@v1.62

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN rm .air.toml
RUN mv relayer.air.toml .air.toml

RUN mkdir -p /app/tmp/air && chmod -R 0777 /app/tmp/air

CMD ["air"]
