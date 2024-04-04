namespace go user

include "base.thrift"

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
    1:i64 user_id,
    2:i64 action_type,
    3:optional string totp,
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
}

struct InfoRequest {
    1:i64 user_id,
}

struct InfoResponse {
    1: base.BaseResp base,
    2: optional User user,
}

struct AvatarRequest{
    1:i64 user_id,
    2:binary avatar_file,
}
struct AvatarResponse{
    1: base.BaseResp base,
    2: optional User user,
}
service UserHandler {
    RegisterResponse Register(1: RegisterRequest req)(api.post="/bibi/user/register/"),
    LoginResponse Login(1: LoginRequest req)(api.post="/bibi/user/login/"),
    InfoResponse Info(1: InfoRequest req)(api.get="/bibi/user/"),
    AvatarResponse Avatar(1:AvatarRequest req)(api.put="/bibi/user/avatar/upload"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bibi/user/2fa"),
    Switch2FAResponse Switch2FA(1:Switch2FARequest req)(api.post="/bibi/user/switch2fa"),
}