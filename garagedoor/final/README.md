# Build

```sh
docker build -t doorcheck:v1 .

```

## Run

```sh
docker run -d \
 --name doorcheck \
 --restart=always \
 -v /etc/localtime:/etc/localtime \
-p 3060:3060 \
-e SLACK_WEBHOOK_URL=${SLACK_WEBHOOK_URL} \
--device /dev/gpiomem --device /dev/mem \
doorcheck:v1 \
-c /etc/config.yml
```
