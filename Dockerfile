FROM golang:tip-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o wa-version

FROM alpine:3.23.2 

WORKDIR /app

COPY --from=builder /app/wa-version ./wa-version

CMD ["./wa-version"]