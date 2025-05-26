# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY calc.go .
RUN \
    apk update \
    && go mod init calc \
    && go mod tidy

RUN \
    go build -o calc calc.go

# This is the final image
FROM alpine

WORKDIR /app
COPY --from=builder /app/calc .

# Create non-root user and group
RUN \ 
    addgroup -g 1000 mygroup \
    && adduser -D -u 1000 -G mygroup myuser \
    && chown myuser:mygroup /app/calc \
    && chmod 755 /app/calc

# Set the user to run the application
USER myuser

EXPOSE 8080

CMD ["./calc"]