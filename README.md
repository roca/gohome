# gohome

## TinyGo Repo

[link](github.com/tinygo-org/tinygo/src/machine)


## CamWatcher

This works best with the 'PI Zero 2 w' the  'PI Zero w' has bad performance issues

```bash
docker build -t camwatcher:v1 .

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
