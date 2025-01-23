FROM docker.io/golang:1.23.4 AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o picotempexport

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D picotempexport
WORKDIR /app
COPY --chown=picotempexport --from=builder /app/picotempexport .
EXPOSE 3030
CMD ["/app/picotempexport"]
