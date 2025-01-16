# Project Notes

## Compile command for Rasberry Pi 5

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o restapi
```

## Build, Tag and Run docker image

```sh
docker build -t restapi:v1 .
docker run -d -p 4000:4000 --name restapi-v1 restapi:v1
```
