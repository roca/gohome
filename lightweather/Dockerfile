FROM docker.io/golang:1.23.4 AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o lightweather lightweather.go config.go

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D lightweather
WORKDIR /app
COPY ./config.yml /etc/config.yml
COPY --chown=lightweather --from=builder /app/lightweather .
EXPOSE 3040
ENTRYPOINT ["/app/lightweather"]
