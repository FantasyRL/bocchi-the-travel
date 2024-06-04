namespace go party

include"base.thrift"


struct CreatePartyRequest{
    1:i64 founder_id,
    2:string title,
    3:string content,
    4:i64 type,
    5: string province,
    6: string city,
    7:string start_time,
    8:string end_time,
}

struct CreatePartyResponse{
    1:base.BaseResp base,
    2:optional base.Party party,
}

struct GetPartyInfoRequest{
    1:i64 party_id,
}

struct GetPartyInfoResponse{
    1:base.BaseResp base,
    2:optional base.Party party,
}


struct JoinPartyRequest{
    1:i64 party_id,
    2:i64 member_id,
}

struct JoinPartyResponse{
    1:base.BaseResp base,
}

struct ApplyListRequest{
    1:i64 party_id,
    2:i64 page_num,
    3:i64 user_id,
}

struct ApplyListResponse{
    1:base.BaseResp base,
    2:optional i64 applicant_count,
    3:optional list<base.User> applicant_list,
}

struct PermitJoinRequest{
    1:i64 party_id,
    2:i64 member_id,
    3:i64 user_id,
}

struct PermitJoinResponse{
    1:base.BaseResp base,
}

struct GetPartyMembersRequest{
    1:i64 party_id,
    2:i64 page_num,
}

struct GetPartyMembersResponse{
    1:base.BaseResp base,
    2:optional i64 member_count,
    3:optional list<base.User> member_list,
}

struct SearchPartyRequest{
    1:optional string content,
    2:optional i64 party_type,
    3:optional string province,
    4:optional string city,
    5:optional i64 start_time_duration,
    6:optional i64 search_type,
    7:i64 page_num,
}

struct SearchPartyResponse{
    1:base.BaseResp base,
    2:optional i64 party_count,
    3:optional list<base.Party> party_list,
}

struct GetMyPartiesRequest{
    1:i64 user_id,
    2:i64 page_num,
}

struct GetMyPartiesResponse{
    1:base.BaseResp base,
    2:optional i64 party_count,
    3:optional list<base.Party> party_list,
}

service PartyHandler{
    CreatePartyResponse CreateParty(1:CreatePartyRequest req)(api.post="/bocchi/party/create"),
    JoinPartyResponse JoinParty(1:JoinPartyRequest req)(api.get="/bocchi/party/apply"),
    ApplyListResponse ApplyList(1:ApplyListRequest req)(api.get="/bocchi/party/apply/list"),
    PermitJoinResponse PermitJoin(1:PermitJoinRequest req)(api.get="/bocchi/party/apply/permit"),
    GetPartyInfoResponse GetPartyInfo(1:GetPartyInfoRequest req)(api.get="/bocchi/party/get"),
    GetPartyMembersResponse GetPartyMembers(1:GetPartyMembersRequest req)(api.get="/bocchi/party/members"),
    SearchPartyResponse SearchParty(1:SearchPartyRequest req)(api.post="/bocchi/party/search"),
    GetMyPartiesResponse GetMyParties(1:GetMyPartiesRequest req)(api.get="/bocchi/party/party/my"),
}