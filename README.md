# A7 和你一起旅行

![project.jpg](project.jpg)

## slowly start

(还未研究容器启动rpc，私密马赛)
```bash
mv config/config-example.yaml config/config.yaml
rm -r config/config-example.yaml
docker-compose up -d
```

```bash
go run bocchi/api
```
```bash
go run bocchi/rpc/user
```
