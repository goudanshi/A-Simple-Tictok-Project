package util

import "time"

const FEED_LIMIT = 30

const USER_ID = "user_id"

const LOCAL_HOST = "http://127.0.0.1:8100"
const GET_VIDEO_PATH = "/douyin/feed/get"

// favorite操作action type
const EXEC_FAVORITE = 1
const CANCEL_FAVORITE = 2

// redis相关参数
const REDIS_ADDRESS = ""
const REDIS_PASSWORD = ""
const REDIS_DB = 0
const REDIS_USERNAME = ""
const REDIS_COMMON_KEY_EXPAIR = 10 * time.Second

// comment操作action type
const PUBLISH_COMMENT = 1
const DELETE_COMMENT = 2

// 关系操作action type
const NEW_RELATION = 1
const DELETE_RELATION = 2

// minIO相关参数
const MINIO_ACCESS_KEY = "admin"
const MINIO_SECRET_ACCESS_KEY = "123456789"
const MINIO_END_POINT = "127.0.0.1:9000/"
const MINIO_BUCKET = "douyinbalute"
