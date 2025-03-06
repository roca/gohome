# gohome

## TinyGo Repo

[link](github.com/tinygo-org/tinygo/src/machine)


## Motion Sensor

This works best with the 'PI Zero 2 w' the  'PI Zero w' has bad performance

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o motion-test main.go
```
