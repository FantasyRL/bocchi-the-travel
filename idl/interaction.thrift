namespace go interaction
include"base.thrift"

//
//struct LikeActionRequest{
//    1:optional i64 video_id,
//    2:optional i64 comment_id,
//    3:i64 action_type,
//    4:i64 user_id,
//}
//
//struct LikeActionResponse{
//    1:base.BaseResp base,
//}
//
//struct LikeListRequest{
//    1:i64 page_num,
//    2:i64 user_id,
//}
//
//struct LikeListResponse{
//    1:base.BaseResp base,
//    2:optional i64 video_count,
//    3:optional list<base.Video> video_list,
//}

struct CommentCreateRequest{
    1:required i64 poi_id,
//    2:optional i64 parent_id,
    2:string content,
    3:i64 user_id,
}

struct CommentCreateResponse{
    1:base.BaseResp base,
}

struct CommentDeleteRequest{
    1:i64 id,
    2:i64 user_id,
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
    3:optional list<base.Comment> comment_list,
}

//内部rpc
//struct GetLikesCountByVideoIdListRequest{
//    1:list<i64> video_id,
//}
//
//struct GetLikesCountByVideoIdListResponse{
//    1:base.BaseResp base,
//    2:list<i64> like_count_list,
//}
//
//struct GetIsLikeByVideoIdListRequest{
//    1:list<i64> video_id,
//    2:i64 user_id,
//}
//
//struct GetIsLikeByVideoIdListResponse{
//    1:base.BaseResp base,
//    2:list<i64> is_like_list,
//}

service InteractionHandler{
//    LikeActionResponse LikeAction(1:LikeActionRequest req),
//    LikeListResponse LikeList(1:LikeListRequest req),
    CommentCreateResponse CommentCreate(1:CommentCreateRequest req),
    CommentDeleteResponse CommentDelete(1:CommentDeleteRequest req),
//    GetLikesCountByVideoIdListResponse GetLikesCountByVideoIdList(1:GetLikesCountByVideoIdListRequest req),
//    GetIsLikeByVideoIdListResponse GetIsLikeByVideoIdList(1:GetIsLikeByVideoIdListRequest req),
    CommentListResponse CommentList(1:CommentListRequest req),
}