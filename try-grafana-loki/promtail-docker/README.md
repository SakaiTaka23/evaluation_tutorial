- GETメソッドの一覧取得

```
{job="containerlogs",stream="stdout"} |= "method=GET"
```

```
sum(count_over_time({job="containerlogs", stream="stdout"}[1m])) by (method)
sum by (method) (rate({job="containerlogs", stream="stdout"}[1m]))
```

- メソッドごとのレート
```
sum by (method) (rate({job="containerlogs", stream="stdout"} | json | __error__=""[1m]))
```

- フォーマット
```
{job="containerlogs",stream="stdout"} |= "host" | json | line_format "{{.host}} {{.method}} {{.status}}"
```

## レジェンド周りの設定
- 「.」はいらさなそう

例
```
{{method}}
```

# Grafanaデータ永続化
- grafana.dbに全て入ってそう

