FROM docker.io/golang:1.23.6 AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o camwatcher camwatcher.go

FROM balenalib/rpi-raspbian:latest
RUN apt update -y && apt upgrade -y
RUN apt install libcamera-tools libcamera-apps-lite -y
RUN apt install libcap-dev libcamera-dev -y
RUN apt install libatlas-base-dev libopenjp2-7 libkms++-dev libfmt-dev libdrm-dev -y

RUN mkdir /app 
WORKDIR /app
COPY --from=builder /app/camwatcher .
RUN chmod -R 755 /app
ENTRYPOINT ["/app/camwatcher"]
