# 起動

```
docker run \
    -p 9090:9090 \
    -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```

# 参考

https://prometheus.io/docs/introduction/first_steps/
