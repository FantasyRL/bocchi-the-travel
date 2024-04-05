package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var (
	OSS           *oss
	Mysql         *mySQL
	Redis         *redis
	Etcd          *etcd
	Server        *server
	Service       *service
	RabbitMQ      *rabbitMQ
	Sender        *email
	ElasticSearch *elasticsearch

	runtimeViper = viper.New()
)

func Init(service string) {
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("yaml")
	runtimeViper.AddConfigPath("./config")
	if err := runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found")
		} else {
			log.Println("config file was found but another error was produced")
		}
	}

	configMapping(service)

}

func configMapping(serviceName string) {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	Server = &c.Server
	Server.Secret = []byte(runtimeViper.GetString("server.jwt-secret"))

	Mysql = &c.MySQL
	RabbitMQ = &c.RabbitMQ
	Etcd = &c.Etcd
	Redis = &c.Redis
	OSS = &c.OSS
	Sender = &c.Email
	ElasticSearch = &c.ElasticSearch

	addrList := runtimeViper.GetStringSlice("services." + serviceName + ".addr")
	Service = &service{
		Name:     runtimeViper.GetString("services." + serviceName + ".name"),
		AddrList: addrList,
		LB:       runtimeViper.GetBool("services." + serviceName + ".load-balance"), //todo:不知道是啥你也ctrl c？
	}

}
