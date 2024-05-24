namespace go base
//公共数据结构
struct BaseResp {
    1: i64 code
    2: string msg
}

struct User {
    1: i64 id,
    2: string name,
    3: string email,
    4: string avatar,
    5:string signature,
}

//struct Route{
//    1:string start_location,
//    2:string end_location,
//    4:double estimated_time,
//    5:double distance,
//    6:string estimated_departure_time,
//    7:string estimated_arrival_time,
//}

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
    12:string rectangle,
    13:list<Itinerary> itineraries,
}

struct Comment {
    1: i64 id,
    2: i64 poi_id,
    4: User user,
    5: string content,
    6: string publish_time,
}

