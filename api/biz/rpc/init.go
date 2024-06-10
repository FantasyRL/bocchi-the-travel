package rpc

import (
	"bocchi/kitex_gen/interaction/interactionhandler"
	"bocchi/kitex_gen/itinerary/itineraryhandler"
	"bocchi/kitex_gen/party/partyhandler"
	"bocchi/kitex_gen/trust/trusthandler"
	"bocchi/kitex_gen/user/userhandler"
)

var (
	userClient        userhandler.Client
	partyClient       partyhandler.Client
	itineraryClient   itineraryhandler.Client
	interactionClient interactionhandler.Client
	trustClient       trusthandler.Client
)

func Init() {
	InitUserRPC()
	InitPartyRPC()
	InitItineraryRPC()
	InitInteractionRPC()
	InitTrustRPC()
}
