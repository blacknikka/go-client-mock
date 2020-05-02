## Grafanaサンプル
- influxDBを扱う例

## GrafanaのInfluxDBプラグインのメモ
### docs
- https://docs.influxdata.com/influxdb/v1.7/tools/api/

### log
- journaldによって管理されている

- 最新のログを見る方法
```bash
sudo journalctl -u influxdb.service -f -n 30
```

- Grafana Datasource PluginからInfluxDBへのリクエストログ
```
May 02 02:07:47 ip-172-31-41-182 influxd[2388]: [httpd] 172.25.0.1, 172.25.0.1,110.132.172.142 - telegraf [02/May/2020:02:07:47 +0000] "GET /query?db=telegraf&epoch=ms&q=SELECT+mean%28%22usage_idle%22%29+FROM+%22cpu%22+WHERE+time+%3E%3D+now%28%29+-+15m+GROUP+BY+time%2810s%29+fill%28null%29 HTTP/1.1" 200 745 "-" "Grafana/6.7.3" b7953cba-8c19-11ea-b1bc-0604e056be56 3686
```

- 上記リクエストにより、以下のクエリが発行される
```
SELECT mean("usage_idle") FROM "cpu" WHERE time >= now() - 15m GROUP BY time(10s) fill(null)
```
また、他のクエリパラメータとして、`db=telegraf`,`epoch=ms`が付与される

#### `epoch=ms`について

タイムスタンプの精度に関するオプションのよう。

|option||説明|
|--|--|--|
|epoch=[ns,u,µ,ms,s,m,h]|Optional|Returns epoch timestamps with the specified precision. By default, InfluxDB returns timestamps in RFC3339 format with nanosecond precision. Both u and µ indicate microseconds.|

