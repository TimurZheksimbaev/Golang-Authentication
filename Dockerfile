FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 9998

RUN go build -o main ./cmd/main.go

CMD ["./main"]



