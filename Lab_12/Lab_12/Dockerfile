FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .


ARG APP_DIR
RUN CGO_ENABLED=0 GOOS=linux go build -o /service ./${APP_DIR}

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /service .
CMD ["./service"]