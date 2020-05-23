FROM golang:1.14 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ipfsadd.go .

RUN GOOS=linux go build -o ipfsadd

FROM debian:stable-slim

COPY --from=builder /app/ipfsadd .

ENTRYPOINT ["./ipfsadd"]
