FROM docker.io/golang:1.23.6 AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o camwatcher

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D camwatcher
WORKDIR /app
COPY --chown=camwatcher --from=builder /app/camwatcher .
ENTRYPOINT ["/app/camwatcher"]
