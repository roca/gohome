# Light Weather App

## Docker build

```sh
docker build -t lightweather:v1 .
```

## Docker Run

```sh`
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
