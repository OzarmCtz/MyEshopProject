# backend/Dockerfile
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Final stage: image avec certificats CA
FROM debian:bookworm-slim
WORKDIR /app

# AJOUTEZ CES LIGNES pour installer les certificats CA
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY appli/datasources/firebase/firebase-service-account-key.json ./appli/datasources/firebase/
COPY .env .
EXPOSE 8081
CMD ["./main"]