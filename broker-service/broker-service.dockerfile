# base go image
FROM golang:1.18-alpine as builder

# make directory for the app
RUN mkdir /app

# copy to app
COPY . /app

WORKDIR /app

# build the app
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest

# copy the binary from builder
RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD [ "/app/brokerApp" ]