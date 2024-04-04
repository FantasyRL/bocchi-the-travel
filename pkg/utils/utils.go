package utils

import (
	"bocchi/config"
	"net"
	"strings"
)

func InitMysqlDSN() string {
	return strings.Join([]string{config.Mysql.User, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=utf8mb4&parseTime=True"}, "")

}

func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)

	if err != nil {
		return false
	}

	l.Close()

	return true
}
