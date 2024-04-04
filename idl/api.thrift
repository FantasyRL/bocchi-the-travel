namespace go api
include"base.thrift"

//user
struct User {
    1: i64 id,
    2: string name,
    3: string email,
    4: i64 follow_count,
    5: i64 follower_count,
    6: bool is_follow,
    7: string avatar,
    8: i64 video_count,
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
service UserHandler {
    RegisterResponse Register(1: RegisterRequest req)(api.post="/bibi/user/register/"),
    LoginResponse Login(1: LoginRequest req)(api.post="/bibi/user/login/"),
    InfoResponse Info(1: InfoRequest req)(api.get="/bibi/user/info"),
    AvatarResponse Avatar(1:AvatarRequest req)(api.put="/bibi/user/avatar/upload"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bibi/user/2fa"),
    Switch2FAResponse Switch2FA(1:Switch2FARequest req)(api.post="/bibi/user/switch2fa"),
    GetAccessTokenResponse GetAccessToken(1:GetAccessTokenRequest req)(api.get="/bibi/access_token/get")
}

//video
struct Video{
    1:i64 id,
    2:string title,
    3:User author,
    4:i64 uid,
    5:string play_url,
    6:string cover_url,
    7:i64 like_count,
    8:i64 comment_count,
    9:i64 is_like,
    10:string publish_time,
}

struct PutVideoRequest{
    1:binary video_file,
    2:string title,
    3:binary cover,
}

struct PutVideoResponse{
    1:base.BaseResp base,
}

struct ListUserVideoRequest{
    1:i64 page_num,
}

struct ListUserVideoResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<Video> video_list,
}

struct SearchVideoRequest{
    1:string param,
    2:i64 page_num,
}

struct SearchVideoResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<Video> video_list,
}

struct HotVideoRequest{
}

struct HotVideoResponse{
    1:base.BaseResp base,
    2:optional list<Video> video_list,
}

service VideoHandler{
    PutVideoResponse PutVideo(1:PutVideoRequest req)(api.post="/bibi/video/upload"),
    ListUserVideoResponse ListVideo(1:ListUserVideoRequest req)(api.get="/bibi/video/published"),
    SearchVideoResponse SearchVideo(1:SearchVideoRequest req)(api.post="/bibi/video/search"),
    HotVideoResponse HotVideo(1:HotVideoRequest req)(api.get="/bibi/video/hot"),
}