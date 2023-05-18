# Base image for golang
FROM golang:1.18-alpine as builder

RUN mkdir /app

# Copy the auth directory
COPY . /app

# Set working dir
WORKDIR /app

RUN CGO_ENABLED=0 go build -o authService ./cmd
# Set execute permission
RUN chmod +x /app/authService

# build a docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authService /app

CMD ["/app/authService"]