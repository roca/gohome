# Prometheus

## Restart

```sh
 docker compose down
 docker volume rm prometheus_prom_data
 docker compose up -d
 docker cp sd_node01.yml prometheus-prometheus-1:/prometheus
 ```
