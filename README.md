# A7 和你一起旅行

![project.jpg](project.jpg)

## 本机裸启动(port:10001)

```bash
#请填写好config-example配置
mv config/config-example.yaml config/config.yaml
make env-up
make start-all
```

[//]: # (## 容器部署)

[//]: # (已写好makefile，打包后直接进入容器，配置好config的端口后运行单个服务即可)

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
[tree.txt](./tree.txt)
```bash
treer -e tree.txt -i "/.idea|.git|data|codeql|(.*)?elasticsearch|(.*)?kibana|web|CNAME/" 
```

