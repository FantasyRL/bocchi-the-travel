package rpc

import (
	"bocchi/kitex_gen/itinerary/itineraryhandler"
	"bocchi/kitex_gen/party/partyhandler"
	"bocchi/kitex_gen/user/userhandler"
)

var (
	userClient      userhandler.Client
	partyClient     partyhandler.Client
	itineraryClient itineraryhandler.Client
)

func Init() {
	InitUserRPC()
	InitPartyRPC()
	InitItineraryRPC()
}
