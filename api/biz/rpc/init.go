package rpc

import (
	"bocchi/kitex_gen/party/partyhandler"
	"bocchi/kitex_gen/user/userhandler"
)

var (
	userClient  userhandler.Client
	partyClient partyhandler.Client
)

func Init() {
	InitUserRPC()
	InitPartyRPC()
}
