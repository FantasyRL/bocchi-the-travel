# A7 和你一起旅行

![project.jpg](project.jpg)

## 本机裸启动(port:10001)

```bash
mv config/config-example.yaml config/config.yaml
make env-up
make start-all
```

## 日志监测
启动后在[kibana](http://127.0.0.1:5601/)中可查看日志


## 技术栈
* api:hertz(thrift)
* db:gorm
* cache:redis
* rpc:kitex(thrift)+etcd
* log:elasticSearch+logrus+kibana

## Page
https://pr.xiey.work/

## tree
```bash
treer -e tree.txt -i "/.idea|.git|data|codeql|(.*)?elasticsearch|(.*)?kibana|web|CNAME/" 
```

