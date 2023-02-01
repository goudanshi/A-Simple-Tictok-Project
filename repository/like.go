package repository

import (
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	. "github.com/A-Simple-Tictok-Project/douyin/utils"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Like struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	UserId    int64          `json:"user_id" gorm:"column:user_id"`
	VideoId   int64          `json:"video_id" gorm:"column:video_id"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`
}

func (l *Like) TableName() string {
	return constants.LikeTableName
}

type LikeDao struct{}

var (
	likeDao  *LikeDao
	likeOnce sync.Once
)

func NewLikeDaoInstance() *LikeDao {
	likeOnce.Do(func() {
		likeDao = &LikeDao{}
	})
	return likeDao
}

// CreateLike 点赞
func (*LikeDao) CreateLike(like *Like) error {
	// 开事务执行点赞操作
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(like).Error; err != nil {
			return err
		}
		if err := tx.Table("video").Where("id = ?", like.VideoId).Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})

}

// DeleteLike 取消点赞
func (*LikeDao) DeleteLike(like *Like) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? and video_id = ?", like.UserId, like.VideoId).Delete(like).Error; err != nil {
			Logger.Info("insert like err: " + err.Error())
			return err
		}
		if err := tx.Table("video").Where("id = ?", like.VideoId).Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// QueryLikeByUserId 喜欢列表
func (*LikeDao) QueryLikeByUserId(userId int64) ([]*Like, error) {
	likes := make([]*Like, 0)
	if err := db.Where("user_id = ?", userId).Find(&likes).Error; err != nil {
		Logger.Error("query like by user id err: " + err.Error())
		return likes, err
	}
	return likes, nil
}
