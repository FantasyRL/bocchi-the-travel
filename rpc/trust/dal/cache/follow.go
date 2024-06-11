package cache

import (
	"bocchi/pkg/constants"
	"bocchi/rpc/trust/dal/db"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func IsFollowedCacheExist(ctx context.Context, followedId int64) (bool, error) {
	ok, err := rFollow.Exists(ctx, i64ToStr(followedId)+constants.FollowerSuffix).Result()
	if err != nil {
		//错误处理返回啥都一样
		return false, err
	}
	if ok > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func IsUserFollowExist(ctx context.Context, followerId int64, followedId int64) (bool, error) {
	return rFollow.SIsMember(ctx, i64ToStr(followedId)+constants.FollowerSuffix, i64ToStr(followerId)).Result()
}

func SetFollowerList(ctx context.Context, uid int64, followList []db.Follow) (err error) {
	tx := rFollow.TxPipeline()
	for _, follow := range followList {
		err = tx.SAdd(ctx, i64ToStr(uid)+constants.FollowerSuffix, i64ToStr(follow.Uid)).Err()
	}
	if err != nil {
		return
	}
	if err = tx.Expire(ctx, i64ToStr(uid)+constants.FollowerSuffix, constants.FollowExpTime).Err(); err != nil {
		return
	}
	_, err = tx.Exec(ctx)
	return
}

func SetFollowingList(ctx context.Context, uid int64, followList *[]db.Follow) (err error) {
	tx := rFollow.TxPipeline()
	for _, follow := range *followList {
		err = tx.SAdd(ctx, i64ToStr(follow.FollowedId)+constants.FollowerSuffix, i64ToStr(uid)).Err()
	}
	if err != nil {
		return
	}
	for _, follow := range *followList {
		err = tx.Expire(ctx, i64ToStr(follow.FollowedId)+constants.FollowerSuffix, constants.FollowExpTime).Err()
		if err != nil {
			return
		}
	}
	_, err = tx.Exec(ctx)
	return err
}

func SetFriendList(ctx context.Context, uid int64, followList *[]db.Follow) (err error) {
	tx := rFollow.TxPipeline()
	for _, follow := range *followList {
		err = tx.SAdd(ctx, i64ToStr(uid)+constants.FollowerSuffix, i64ToStr(follow.Uid)).Err()
		if err != nil {
			return
		}
		err = tx.SAdd(ctx, i64ToStr(follow.FollowedId)+constants.FollowerSuffix, i64ToStr(uid)).Err()
		if err != nil {
			return
		}
	}

	if err = tx.Expire(ctx, i64ToStr(uid)+constants.FollowerSuffix, constants.FollowExpTime).Err(); err != nil {
		return
	}
	for _, follow := range *followList {
		err = tx.Expire(ctx, i64ToStr(follow.FollowedId)+constants.FollowerSuffix, constants.FollowExpTime).Err()
		if err != nil {
			return
		}
	}
	_, err = tx.Exec(ctx)
	return
}

func SetFollowerCount(ctx context.Context, uid int64, count int64) (err error) {
	err = rFollow.ZAdd(ctx, constants.FollowerCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(uid),
	}).Err()
	if err != nil {
		return
	}
	err = rFollow.Expire(ctx, constants.FollowerCountZset, constants.FollowExpTime).Err()
	return
}

func SetFollowingCount(ctx context.Context, uid int64, count int64) (err error) {
	err = rFollow.ZAdd(ctx, constants.FollowingCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(uid),
	}).Err()
	if err != nil {
		return
	}
	err = rFollow.Expire(ctx, constants.FollowingCountZset, constants.FollowExpTime).Err()
	return
}

func SetFriendCount(ctx context.Context, uid int64, count int64) (err error) {
	err = rFollow.ZAdd(ctx, constants.FriendCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(uid),
	}).Err()
	if err != nil {
		return err
	}
	err = rFollow.Expire(ctx, constants.FriendCountZset, constants.FollowExpTime).Err()
	return
}

func AddFollower(ctx context.Context, followerId int64, followedId int64) (err error) {
	tx := rFollow.TxPipeline()
	if err = tx.SAdd(ctx, i64ToStr(followedId)+constants.FollowerSuffix, i64ToStr(followerId)).Err(); err != nil {
		return err
	}
	if err = tx.ZIncrBy(ctx, constants.FollowerCountZset, 1, i64ToStr(followedId)).Err(); err != nil {
		return err
	}
	if err = tx.Expire(ctx, i64ToStr(followedId)+constants.FollowerSuffix, constants.FollowExpTime).Err(); err != nil {
		return
	}
	if err = tx.Expire(ctx, constants.FollowerCountZset, constants.FollowExpTime).Err(); err != nil {
		return
	}
	_, err = tx.Exec(ctx)
	return
}

func DelFollower(ctx context.Context, followerId int64, followedId int64) (err error) {
	tx := rFollow.TxPipeline()
	if err = tx.SRem(ctx, i64ToStr(followedId)+constants.FollowerSuffix, i64ToStr(followerId)).Err(); err != nil {
		return err
	}
	if err = tx.ZIncrBy(ctx, constants.FollowerCountZset, -1, i64ToStr(followedId)).Err(); err != nil {
		return err
	}
	if err = tx.Expire(ctx, i64ToStr(followedId)+constants.FollowerSuffix, constants.FollowExpTime).Err(); err != nil {
		return
	}
	if err = tx.Expire(ctx, constants.FollowerCountZset, constants.FollowExpTime).Err(); err != nil {
		return
	}
	_, err = tx.Exec(ctx)
	return
}

func AddFollowing(ctx context.Context, followerId int64) (err error) {
	if err = rFollow.ZIncrBy(ctx, constants.FollowingCountZset, 1, i64ToStr(followerId)).Err(); err != nil {
		return
	}
	if err = rFollow.Expire(ctx, constants.FollowingCountZset, constants.FollowExpTime).Err(); err != nil {
		return
	}
	return nil
}

func DelFollowing(ctx context.Context, followerId int64) (err error) {
	if err = rFollow.ZIncrBy(ctx, constants.FollowingCountZset, 1, i64ToStr(followerId)).Err(); err != nil {
		return
	}
	if err = rFollow.Expire(ctx, constants.FollowingCountZset, constants.FollowExpTime).Err(); err != nil {
		return
	}
	return nil
}

func GetFollower(ctx context.Context, uid int64) ([]db.Follow, error) {
	vals, err := rFollow.SMembers(ctx, i64ToStr(uid)+constants.FollowerSuffix).Result()
	if err != nil {
		return nil, err
	}
	var followerList []db.Follow
	for _, follower := range vals {
		followerId, _ := strconv.ParseInt(follower, 10, 64)
		followerList = append(followerList, db.Follow{
			Uid:        followerId,
			FollowedId: uid,
		})
	}
	return followerList, nil
}

func GetFollowerCount(ctx context.Context, uid int64) (bool, int64, error) {
	v, err := rFollow.ZScore(ctx, constants.FollowerCountZset, i64ToStr(uid)).Result()
	if errors.Is(err, redis.Nil) { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}

func GetFollowingCount(ctx context.Context, uid int64) (bool, int64, error) {
	v, err := rFollow.ZScore(ctx, constants.FollowingCountZset, i64ToStr(uid)).Result()
	if errors.Is(err, redis.Nil) { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}
