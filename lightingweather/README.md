#  Light Weather App

## Docker Run:

```sh
docker run -d \
--name lightweather \
--restart=always \
-v ${pwd}/config.yml:/etc/config.yml \
-p 3040:3040 \
-e  \
ZIPCODE=${ZIPCODE} \
COUNTRY_CODE=${COUNTRY_CODE} \
OWM_API_KEY=${OWM_API_KEY} \
HUE_ID=${HUE_ID} \
--net=prometheus_prom_net \
lightweather:v1 \
-c /etc/config.yml
```
