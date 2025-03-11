# gohome

## TinyGo Repo

[link: https://github.com/tinygo-org/tinygo](https://github.com/tinygo-org/tinygo)

## cyw43439 Repo@

[link: https://github.com/soypat/cyw43439](https://github.com/soypat/cyw43439)

## Temperature Monitor

```sh
tinygo build -target=pico-w -opt=1 -stack-size=8kb -size=short -serial=uart -o main02.uf2 main.go
```


## Garage Door

## Light Weather

```sh
bash
export ZIPCODE
export COUNTRY_CODE
export OWM_API_KEY
export HUE_ID
export HUE_IP_ADDRESS
```

## Light Weather App

### Docker build

```sh
docker build -t lightweather:v1 .
```

### Docker Run

```sh
docker run -d \
--name lightweather \
--restart=always \
-p 3040:3040 \
-e  ZIPCODE=${ZIPCODE} \
-e COUNTRY_CODE=${COUNTRY_CODE} \
-e OWM_API_KEY=${OWM_API_KEY} \
-e HUE_ID=${HUE_ID} \
-e HUE_IP_ADDRESS=${HUE_IP_ADDRESS} \
--net=prometheus_prom_net \
lightweather:v1 \
-c /etc/config.yml
```

## Cam Watcher

```sh
export SLACK_BOT_TOKEN=
export SLACK_WEBHOOK_URL
export CHANNEL_ID
export CHANNEL_NAME
```

This works best with the 'PI Zero 2 w' the  'PI Zero w' has bad performance issues

### Docker build

```sh
docker build -t camwatcher:v1 .
```

### Docker run

```sh
docker run -d \
-v /run/udev/:/run/udev:ro \
-v /dev/video0:/dev/video0 \
--privileged \
--name camwatcher-v1 \
--env SLACK_BOT_TOKEN=$SLACK_BOT_TOKEN \
--env SLACK_WEBHOOK_URL=$SLACK_WEBHOOK_URL \
--env CHANNEL_ID=$CHANNEL_ID \
--env CHANNEL_NAME=$CHANNEL_NAME \
--restart=always \
camwatcher:v1
```
