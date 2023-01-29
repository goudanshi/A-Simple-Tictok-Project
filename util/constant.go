package util

import "time"

const FEED_LIMIT = 30

const USER_ID = "user_id"

const LOCAL_HOST = "http://192.168.31.20:8100"
const OSS_SHARE_HOST = "http://192.168.31.20:9000/douyinbalute/"
const GET_VIDEO_PATH = "/douyin/feed/get"
const GET_COVER_PATH = "/douyin/feed/get/cover"

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
const MINIO_ACCESS_KEY = "minioadmin"
const MINIO_SECRET_ACCESS_KEY = "minioadmin"
const MINIO_END_POINT = "127.0.0.1:9000/"
const MINIO_BUCKET = "douyinbalute"
