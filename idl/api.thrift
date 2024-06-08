namespace go api

include"base.thrift"

//user
struct User {
    1: i64 id,
    2: string name,
    3: string email,
    4: string avatar,
    5:string signature,
    6:bool is_follow,
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

struct VerifyAccessTokenRequest{

}

struct VerifyAccessTokenResponse{
    1:base.BaseResp base,
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
    1:string signature,
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
    VerifyAccessTokenResponse VerifyAccessToken(1:VerifyAccessTokenRequest req)(api.get="/bocchi/access_token/verify"),
}

//party

struct Party{
    1: i64 id,
    2: i64 founder_id,
    3: string title,
    4: string content,
    5: string type,
    6: string province,
    7: string city,
    8:string start_time,
    9:string end_time,
    10:i64 member_count,
    11:i64 status,
    12:string rectangle,
}

struct CreatePartyRequest{
    1:string title,
    2:string content,
    3:string type,
    4: string province,
    5: string city,
    6:string start_time,
    7:string end_time,
}

struct CreatePartyResponse{
    1:base.BaseResp base,
    2:optional Party party,
}

struct GetPartyInfoRequest{
    1:i64 party_id,
}

struct GetPartyInfoResponse{
    1:base.BaseResp base,
    2:optional Party party,
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
    2:optional string party_type,
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

struct GetMyPartiesRequest{
    1:i64 page_num,
}

struct GetMyPartiesResponse{
    1:base.BaseResp base,
    2:optional i64 party_count,
    3:optional list<Party> party_list,
}


struct ChangePartyStatusRequest{
    1:i64 party_id,
    2:i64 action_type,
}

struct ChangePartyStatusResponse{
    1:base.BaseResp base,
}

struct AddAdminRequest{
    1:i64 party_id,
    2:i64 target_id,
}
struct AddAdminResponse{
    1:base.BaseResp base,
}

struct DeleteAdminRequest{
    1:i64 party_id,
    2:i64 target_id,
}
struct DeleteAdminResponse{
    1:base.BaseResp base,
}

struct DeleteMemberRequest{
    1:i64 party_id,
    2:i64 target_id,
}
struct DeleteMemberResponse{
    1:base.BaseResp base,
}
service PartyHandler{
    DeleteMemberResponse DeleteMember(1:DeleteMemberRequest req)(api.get="/bocchi/party/member/delete"),
    AddAdminResponse AddAdmin(1:AddAdminRequest req)(api.get="/bocchi/party/admin/create"),
    DeleteAdminResponse DeleteAdmin(1:DeleteAdminRequest req)(api.get="/bocchi/party/admin/delete"),
    CreatePartyResponse CreateParty(1:CreatePartyRequest req)(api.post="/bocchi/party/create"),
    JoinPartyResponse JoinParty(1:JoinPartyRequest req)(api.get="/bocchi/party/apply"),
    ApplyListResponse ApplyList(1:ApplyListRequest req)(api.get="/bocchi/party/apply/list"),
    PermitJoinResponse PermitJoin(1:PermitJoinRequest req)(api.get="/bocchi/party/apply/permit"),
    GetPartyInfoResponse GetPartyInfo(1:GetPartyInfoRequest req)(api.get="/bocchi/party/get"),
    GetPartyMembersResponse GetPartyMembers(1:GetPartyMembersRequest req)(api.get="/bocchi/party/members"),
    SearchPartyResponse SearchParty(1:SearchPartyRequest req)(api.post="/bocchi/party/search"),
    GetMyPartiesResponse GetMyParties(1:GetMyPartiesRequest req)(api.get="/bocchi/party/party/my"),
    ChangePartyStatusResponse ChangePartyStatus(1:ChangePartyStatusRequest req)(api.get="/bocchi/party/status"),
}

//itinerary
struct Itinerary{
    1:i64 id,
    2:string title,
    3:i64 founder_id,
    4:i64 action_type,
    5:optional string rectangle,
    6:optional string route_json,
    7:optional string remark,
    8:i64 sequence,
    9:string schedule_start_time,
    10:string schedule_end_time,
    11:i64 party_id,
    12:i64 is_merged,
}

struct CreateItineraryRequest{
    1:string title,
    2:i64 action_type,
    3:optional string rectangle,
    4:optional string route_json,
    5:optional string remark,
    6:string schedule_start_time,
    7:string schedule_end_time,
    8:i64 party_id,
}

struct CreateItineraryResponse{
    1:base.BaseResp base,
    2:optional Itinerary itinerary,
}

struct GetItineraryInfoRequest{
    1:i64 itinerary_id,
}

struct GetItineraryInfoResponse{
    1:base.BaseResp base,
    2:optional Itinerary itinerary,
}


struct ShowPartyItineraryRequest{
    1:i64 party_id,
}

struct ShowPartyItineraryResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<Itinerary> itineraries,
}

struct ChangeSequenceRequest{
    1:i64 party_id,
    2:list<i64> itinerary_id_list,
    3:list<i64> sequence_list,
}

struct ChangeSequenceResponse{
    1:base.BaseResp base,
}

struct MergeItineraryRequest{
    1:i64 party_id,
    2:i64 itinerary_id,
}

struct MergeItineraryResponse{
    1:base.BaseResp base,
}

struct GetMyItinerariesRequest{
    1:i64 party_id,

}

struct GetMyItinerariesResponse{
    1:base.BaseResp base,
    2:optional i64 itienrary_count,
    3:optional list<Itinerary> itinerary_list,
}

struct DeleteItineraryRequest{
    1:i64 itinerary_id,
}

struct DeleteItineraryResponse{
    1:base.BaseResp base,
}

struct ShowItineraryDraftRequest{
    1:i64 party_id,
}

struct ShowItineraryDraftResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<Itinerary> itineraries,
}

service ItineraryHandler{
    CreateItineraryResponse CreateItinerary(1:CreateItineraryRequest req)(api.post="/bocchi/party/itinerary/create"),
    GetItineraryInfoResponse GetItineraryInfo(1:GetItineraryInfoRequest req)(api.get="/bocchi/party/itinerary/info"),
    ShowPartyItineraryResponse ShowPartyItinerary(1:ShowPartyItineraryRequest req)(api.get="/bocchi/party/itinerary/show"),
    ShowItineraryDraftResponse ShowItineraryDraft(1:ShowItineraryDraftRequest req)(api.get="/bocchi/party/itinerary/draft/show"),
    ChangeSequenceResponse ChangeSequence(1:ChangeSequenceRequest req)(api.post="/bocchi/party/itinerary/sequence/change"),
    MergeItineraryResponse MergeItinerary(1:MergeItineraryRequest req)(api.get="/bocchi/party/itinerary/merge"),
    GetMyItinerariesResponse GetMyItineraries(1:GetMyItinerariesRequest req)(api.get="/bocchi/party/itinerary/my"),
    DeleteItineraryResponse DeleteItinerary(1:DeleteItineraryRequest req)(api.get="/bocchi/party/itinerary/delete"),
}

//comment
struct Comment {
    1: i64 id,
    2: i64 poi_id,
    4: User user,
    5: string content,
    6: string publish_time,
}

struct CommentCreateRequest{
    1:required i64 poi_id,
    2:string content,
}

struct CommentCreateResponse{
    1:base.BaseResp base,
    2:optional i64 comment_id,
}

struct CommentDeleteRequest{
    1:i64 comment_id,
}

struct CommentDeleteResponse{
    1:base.BaseResp base,
}

struct CommentListRequest{
    1:i64 poi_id,
    2:i64 page_num,
}

struct CommentListResponse{
    1:base.BaseResp base,
    2:optional i64 comment_count,
    3:optional list<Comment> comment_list,
}

service InteractionHandler{
    CommentCreateResponse CommentCreate(1:CommentCreateRequest req)(api.post="/bocchi/poi/comment/create"),
    CommentDeleteResponse CommentDelete(1:CommentDeleteRequest req)(api.get="/bocchi/poi/comment/delete"),
    CommentListResponse CommentList(1:CommentListRequest req)(api.get="/bocchi/poi/comment/list"),
}

//trust
struct FollowActionRequest{
    1: i64 object_uid,
    2: i64 action_type,
    3: i64 user_id,
}

struct FollowActionResponse{
    1:base.BaseResp base,
}

struct FollowingListRequest{
    1:i64 page_num,
    2:i64 user_id,
}

struct FollowingListResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<User> following_list,
}

struct FollowerListRequest{
    1: i64 page_num,
    2: i64 user_id,
}

struct FollowerListResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<User> follower_list,
}

struct FriendListRequest{
    1: i64 page_num,
}

struct FriendListResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<User> friend_list,
}

struct MarkToOtherRequest{
    1:i64 target_id,
    2:double score,
}

struct MarkToOtherResponse{
    1:base.BaseResp base,
}

struct GetUserScoreRequest{
    1:i64 user_id,
}
struct GetUserScoreResponse{

}

service TrustHandler{
    FollowActionResponse TrustAction(1:FollowActionRequest req)(api.post="/bibi/trust/action"),
    FollowerListResponse FollowerList(1:FollowerListRequest req)(api.get="/bibi/trust/follower"),
    FollowingListResponse FollowingList(1:FollowingListRequest req)(api.get="/bibi/trust/following"),
    MarkToOtherResponse MarkToOther(1:MarkToOtherRequest req)(api.get="/bibi/trust/mark"),
    FriendListResponse TrustEachList(1:FriendListRequest req)(api.post="/bibi/trust/each"),
}