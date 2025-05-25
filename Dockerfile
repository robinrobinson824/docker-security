# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY app.go .
RUN \
    apk update \
    && go mod init app \
    && go mod tidy
# go build -o app means build the app.go file and output the binary to app

RUN \
    go build -o app app.go

# This is the final image
FROM alpine

# Create non-root user and group
RUN \ 
    addgroup -g 1000 mygroup \
    && adduser -D -u 1000 -G mygroup myuser
# addgroup -g 1000 mygroup means add a group with gid 1000
# adduser -D -u 1000 -G mygroup myuser means add a user with uid 1000 and gid 1000 to the group mygroup
# -D means add a system user

WORKDIR /app
COPY --from=builder /app/app .
# /app/app . means copy the app binary from the builder stage to the current working directory
# --from=builder means copy from the builder stage
# --chown=app:app means change the owner of the file to app:app

RUN \ 
    chown myuser:mygroup /app/app \
    && chmod 755 /app/app

USER myuser
EXPOSE 8080
#USER app
CMD ["./app"]