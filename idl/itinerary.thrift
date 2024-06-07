namespace go itinerary
include"base.thrift"

struct CreateItineraryRequest{
    1:i64 founder_id,
    2:string title,
    3:i64 action_type,
    4:optional string rectangle,
    5:optional string route_json,
    6:optional string remark,
    7:string schedule_start_time,
    8:string schedule_end_time,
    9:i64 party_id,
}

struct CreateItineraryResponse{
    1:base.BaseResp base,
    2:optional base.Itinerary itinerary,
}

struct GetItineraryInfoRequest{
    1:i64 itinerary_id,
}

struct GetItineraryInfoResponse{
    1:base.BaseResp base,
    2:optional base.Itinerary itinerary,
}

struct ShowPartyItineraryRequest{
    1:i64 party_id,
}

struct ShowPartyItineraryResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<base.Itinerary> itineraries,
}

struct ChangeSequenceRequest{
    1:list<i64> itinerary_id_list,
    2:list<i64> sequence_list,
    3:i64 user_id,
}

struct ChangeSequenceResponse{
    1:base.BaseResp base,
}

struct MergeItineraryRequest{
    1:i64 party_id,
    2:i64 itinerary_id,
    3:i64 user_id,
}

struct MergeItineraryResponse{
    1:base.BaseResp base,
}

struct GetMyItinerariesRequest{
    1:i64 user_id,
    2:i64 party_id,
}

struct GetMyItinerariesResponse{
    1:base.BaseResp base,
    2:optional i64 itienrary_count,
    3:optional list<base.Itinerary> itinerary_list,
}
//todo:delete itinerary
struct DeleteItineraryRequest{
    1:i64 itinerary_id,
    2:i64 user_id,
}

struct DeleteItineraryResponse{
    1:base.BaseResp base,
}
//todo:itinerary draft

service ItineraryHandler{
    CreateItineraryResponse CreateItinerary(1:CreateItineraryRequest req)(api.post="/bocchi/party/itinerary/create"),
    GetItineraryInfoResponse GetItineraryInfo(1:GetItineraryInfoRequest req)(api.get="/bocchi/party/itinerary/info"),
    ShowPartyItineraryResponse ShowPartyItinerary(1:ShowPartyItineraryRequest req)(api.get="/bocchi/party/itinerary/show"),
    ChangeSequenceResponse ChangeSequence(1:ChangeSequenceRequest req)(api.post="/bocchi/party/itinerary/sequence/change"),
    MergeItineraryResponse MergeItinerary(1:MergeItineraryRequest req)(api.get="/bocchi/party/itinerary/merge"),
    GetMyItinerariesResponse GetMyItineraries(1:GetMyItinerariesRequest req)(api.get="/bocchi/party/itinerary/my"),
    DeleteItineraryResponse DeleteItinerary(1:DeleteItineraryRequest req)(api.get="/bocchi/party/itinerary/delete"),
}