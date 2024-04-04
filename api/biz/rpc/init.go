package rpc

import (
	"bocchi/kitex_gen/user/userhandler"
)

var (
	userClient userhandler.Client
)

func Init() {
	InitUserRPC()
}
