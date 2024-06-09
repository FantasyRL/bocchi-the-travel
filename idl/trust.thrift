namespace go trust

include "base.thrift"

struct FollowActionRequest{
    1: i64 object_uid,
    2: i64 action_type,
    3: i64 user_id,
}

struct FollowActionResponse{
    1:base.BaseResp base,
}

struct FollowingListRequest{
    1: i64 page_num,
    2:i64 user_id,
}

struct FollowingListResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<base.User> following_list,
}

struct FollowerListRequest{
    1: i64 page_num,
    2: i64 user_id,
}

struct FollowerListResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<base.User> follower_list,
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

struct MarkToOtherRequest{
    1:i64 target_id,
    2:double score,
    3:i64 user_id,
}

struct MarkToOtherResponse{
    1:base.BaseResp base,
}

struct GetUserScoreRequest{
    1:i64 user_id,
}
struct GetUserScoreResponse{
    1:base.BaseResp base,
    2:optional double score,
    3:optional i64 count,
}

service TrustHandler{
    FollowActionResponse TrustAction(1:FollowActionRequest req),
    FollowerListResponse FollowerList(1:FollowerListRequest req),
    FollowingListResponse FollowingList(1:FollowingListRequest req),
    MarkToOtherResponse MarkToOther(1:MarkToOtherRequest req),
    FriendListResponse TrustEachList(1:FriendListRequest req),
    GetUserScoreResponse GetUserScore(1:GetUserScoreRequest req)(api.get="/bibi/user/score"),

}