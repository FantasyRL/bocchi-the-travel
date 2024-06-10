namespace go friend

include"base.thrift"

struct ApplyFriendRequest{
    1: i64 object_uid,
    2: i64 action_type,
    3: i64 user_id,
}

struct ApplyFriendResponse{
    1:base.BaseResp base,
}

struct PermitFriendRequest{
    1: i64 object_uid,
    2: i64 action_type,
    3: i64 user_id,
}

struct PermitFriendResponse{
    1:base.BaseResp base,
}

struct FriendListRequest{
    1: i64 page_num,
    2: i64 user_id,
}

struct FriendListResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<base.User> friend_list,
}

struct DeleteFriendRequest{

}

struct DeleteFriendResponse{

}

service FollowHandler{
    ApplyFriendResponse ApplyFriend(1:ApplyFriendRequest req),
    PermitFriendResponse PermitFriend(1:PermitFriendRequest req),
    FriendListResponse FriendList(1:FriendListRequest req),
}