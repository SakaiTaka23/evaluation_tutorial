- ログレベルごとの取得

```
sum by (level) (count_over_time({filename="/var/lib/docker/containers/a19d833869cb41ced627dd42ceac2cddf0b9ef1f138b3287815fda811455346d/a19d833869cb41ced627dd42ceac2cddf0b9ef1f138b3287815fda811455346d-json.log"} | json[1m]))
```

- ログのJSON出力
```
{filename="/var/lib/docker/containers/a19d833869cb41ced627dd42ceac2cddf0b9ef1f138b3287815fda811455346d/a19d833869cb41ced627dd42ceac2cddf0b9ef1f138b3287815fda811455346d-json.log"} | json
```

- フォーマット
```
{job="containerlogs",stream="stdout"} |= "host" | json | line_format "{{.host}} {{.method}} {{.status}}"
```
