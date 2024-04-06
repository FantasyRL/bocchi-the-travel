# A7 和你一起旅行

![project.jpg](project.jpg)

## slowly start

(还未研究容器启动rpc，私密马赛)
```bash
mv config/config-example.yaml config/config.yaml
make env-up
```

```bash
make api
```

```bash
make user
```

```bash
make party
```

## 技术栈
* api:hertz(thrift)
* db:gorm
* cache:redis
* rpc:kitex(thrift)+etcd
* log:elasticSearch+logrus+kibana
