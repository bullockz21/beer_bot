FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY internal/config/.env /internal/config/.env
COPY .env .env

RUN go build -o beer_bot ./cmd/main.go

CMD ["./beer_bot"]
