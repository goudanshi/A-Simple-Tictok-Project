package repository

import (
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	. "github.com/A-Simple-Tictok-Project/douyin/utils"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	ID           int64          `json:"id" gorm:"primarykey"`
	Title        string         `json:"title" gorm:"column:title"`
	PlayURL      string         `json:"play_url" gorm:"column:play_url"`
	CoverURL     string         `json:"cover_url" gorm:"column:cover_url"`
	UserId       int64          `json:"user_id" gorm:"column:user_id"`
	CommentCount int64          `json:"comment_count" gorm:"column:comment_count;default:0"`
	LikeCount    int64          `json:"like_count" gorm:"column:like_count;default:0"`
	CreatedAt    time.Time      `json:"created_at" gorm:"column:create_at"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeleteAt     gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`

	// 外键字段
	Comments []*Comment `json:"comments" gorm:"foreignKey:VideoId;references:ID"`
	Likes    []*Like    `json:"likes" gorm:"foreignKey:VideoId;references:ID"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

type VideoDao struct{}

var (
	videoDao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

// CreateVideo 视频投稿
func (*VideoDao) CreateVideo(video *Video) (int64, error) {
	err := db.Create(video).Error
	if err != nil {
		Logger.Error("insert video err: " + err.Error())
		return 0, err
	}
	return video.ID, nil
}

// QueryVideoByUserId 发布列表
func (*VideoDao) QueryVideoByUserId(userId int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := db.Preload("Comments").Preload("Likes").Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		Logger.Error("find videos by user id err: " + err.Error())
		return videos, err
	}
	return videos, nil
}

// QueryVideoByVideoId 发布列表
func (*VideoDao) QueryVideoByVideoId(videoIds []int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := db.Preload("Comments").Preload("Likes").Where("id in (?)", videoIds).Find(&videos).Error
	if err != nil {
		Logger.Error("find videos by video id err: " + err.Error())
		return videos, err
	}
	return videos, nil
}

// QueryVideoByTime 视频流
func (*VideoDao) QueryVideoByTime(latestTime string) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := db.Where("create_at < ?", latestTime).Limit(constants.VideoLimitCount).Order("create_at desc").Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}
