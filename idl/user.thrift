namespace go user

include "base.thrift"



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
    2: optional base.User user,
}

struct InfoRequest {
    1:i64 user_id,
}

struct InfoResponse {
    1: base.BaseResp base,
    2: optional base.User user,
}

struct AvatarRequest{
    1:i64 user_id,
    2:binary avatar_file,
}
struct AvatarResponse{
    1: base.BaseResp base,
    2: optional base.User user,
}

struct SignatureRequest{
    1:i64 user_id,
    2:string signature,
}

struct SignatureResponse{
    1: base.BaseResp base,
}


//rpc
struct GetUsersRequest{
    1:list<i64> user_id_list,
}

struct GetUsersResponse{
    1:base.BaseResp base,
    2:list<base.User> user_list,
}

service UserHandler {
    RegisterResponse Register(1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    InfoResponse Info(1: InfoRequest req),
    AvatarResponse Avatar(1:AvatarRequest req),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req),
    Switch2FAResponse Switch2FA(1:Switch2FARequest req),
    SignatureResponse Signature(1:SignatureRequest req),

    GetUsersResponse GetUserList(1:GetUsersRequest req),
}