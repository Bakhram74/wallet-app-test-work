FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/app/main.go

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY config.env .
COPY ./migrations ./migrations

CMD [ "/app/main" ]