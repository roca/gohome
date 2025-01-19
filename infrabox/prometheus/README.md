# Prometheus

## Restart

```sh
 docker compose down
 docker volume rm prometheus_prom_data
 docker compose up -d
 docker cp sd_node01.yml prometheus-prometheus-1:/prometheus
 dexec prometheus-prometheus-1 ls

 # Grafana

 ```sh
 docker run -d \
 --name=grafana01 \
--restart=always \
--net=prometheus_prom_net \
-p 3000:3000 \
docker.io/grafana/grafana-oss:95.6
```

```

# Grafana

```sh
docker run -d \
--name=grafana01 \
--restart=always \
--net=prometheus_prom_net \
-p 3000:3000 \
docker.io/grafana/grafana-oss:9.5.6
```

