package util

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDRESS,
		Username: REDIS_USERNAME,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})
	return nil
}

func Set(key string, value string, expire int64) error {
	return rdb.Set(ctx, key, value, time.Duration(expire)).Err()
}

func Get(key string) (string, error) {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func Del(key string) error {
	return rdb.Del(ctx, key).Err()
}

/**
redis梳理：
本质其实只有用户表和视频表
点赞和评论本质还是绑定着视频
关系则绑定着用户
key设计：
user:$id
video:$id
comment_list:$video_id
favorite_list:$video_id
follow_list:$user_id
follower_list:$user_id
目前的问题：在列表接口里，需要获取用户信息，但是用户信息不可能一起放在这个key里，不然用户信息变的时候需要删除的缓存太多
但如果不放进去，就意味着查出来的结果需要每条再去单独查一次用户信息，也不可能
可能的解决方法，单独去查用户信息，且不走redis，直接用mysql的in
*/
