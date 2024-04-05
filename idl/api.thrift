namespace go api
include"base.thrift"

//user
struct User {
    1: i64 id,
    2: string name,
    3: string email,
    4: string avatar,
    5:string signature,
}

struct RegisterRequest {
    1: string username,
    2: string email,
    3: string password,
}

struct RegisterResponse {
    1: base.BaseResp base,
    2: optional i64 user_id,
}

//struct OTP2FAReq{
//    1:i64 uid,
//}

//struct OTP2FAResp{
//    1:BaseResp base,
//}

struct Switch2FARequest{
    1:i64 action_type,
    2:optional string totp,
}

struct Switch2FAResponse{
    1:base.BaseResp base,
}

struct LoginRequest {
    1: string username,
    2: string password,
    3: optional string otp,
}

struct LoginResponse {
    1: base.BaseResp base,
    2: optional User user,
    3: optional string access_token,
    4: optional string refresh_token,
}

struct InfoRequest {
    1:i64 user_id,
}

struct GetAccessTokenRequest{

}

struct GetAccessTokenResponse{
    1:base.BaseResp base,
    2:optional string access_token,
}


struct InfoResponse {
    1: base.BaseResp base,
    2: optional User user,
}

struct AvatarRequest{
    1:binary avatar_file,
}
struct AvatarResponse{
    1: base.BaseResp base,
    2: optional User user,
}

struct SignatureRequest{
    1:i64 user_id,
    2:string signature,
}

struct SignatureResponse{
    1: base.BaseResp base,
}

service UserHandler {
    RegisterResponse Register(1: RegisterRequest req)(api.post="/bocchi/user/register/"),
    LoginResponse Login(1: LoginRequest req)(api.post="/bocchi/user/login/"),
    InfoResponse Info(1: InfoRequest req)(api.get="/bocchi/user/info"),
    AvatarResponse Avatar(1:AvatarRequest req)(api.put="/bocchi/user/avatar/upload"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bocchi/user/2fa"),
    Switch2FAResponse Switch2FA(1:Switch2FARequest req)(api.post="/bocchi/user/switch2fa"),
    GetAccessTokenResponse GetAccessToken(1:GetAccessTokenRequest req)(api.get="/bocchi/access_token/get"),
    SignatureResponse Signature(1:SignatureRequest req)(api.post="/bocchi/user/signature"),

}

//party

struct Party{
    1: i64 id,
    2: i64 founder_id,
    3: string title,
    4: string content,
    5: i64 type,
    6: string province,
    7: string city,
    8:string start_time,
    9:string end_time,
    10:i64 member_count,
    11:i64 status,
}

struct CreatePartyRequest{
    1:string title,
    2:string content,
    3:i64 type,
    4: string province,
    5: string city,
    6:string start_time,
    7:string end_time,
}

struct CreatePartyResponse{
    1:base.BaseResp base,
}

struct JoinPartyRequest{
    1:i64 party_id,
}

struct JoinPartyResponse{
    1:base.BaseResp base,
}

struct ApplyListRequest{
    1:i64 party_id,
    2:i64 page_num,
}

struct ApplyListResponse{
    1:base.BaseResp base,
    2:optional i64 applicant_count,
    3:optional list<User> applicant_list,
}

struct PermitJoinRequest{
    1:i64 party_id,
    2:i64 member_id,
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
    3:optional list<User> member_list,
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
    3:optional list<Party> party_list,
}

service PartyHandler{
    CreatePartyResponse CreateParty(1:CreatePartyRequest req)(api.post="/bocchi/party/create"),
    JoinPartyResponse JoinParty(1:JoinPartyRequest req)(api.get="/bocchi/party/apply"),
    ApplyListResponse ApplyList(1:ApplyListRequest req)(api.get="/bocchi/party/apply/list"),
    PermitJoinResponse PermitJoin(1:PermitJoinRequest req)(api.get="/bocchi/party/apply/permit"),
    GetPartyMembersResponse GetPartyMembers(1:GetPartyMembersRequest req)(api.get="/bocchi/party/members"),
    SearchPartyResponse SearchParty(1:SearchPartyRequest req)(api.post="/bocchi/party/search"),

}