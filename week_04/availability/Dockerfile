# Étape 1 : compilation
FROM golang:1.26 AS builder
WORKDIR /app
COPY go.mod .
COPY *.go .
RUN go build -o simulator .

# Étape 2 : image finale
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/simulator .
CMD ["./simulator"]
