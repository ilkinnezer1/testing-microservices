# Get base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerService ./cmd

RUN chmod +x /app/brokerService


# build a docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerService /app

CMD [ "/app/brokerService" ]