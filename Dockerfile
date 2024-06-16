FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/main .

COPY app.env .

EXPOSE 8081

CMD ["./main"]



