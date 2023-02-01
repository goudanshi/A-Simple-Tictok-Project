package repository

import (
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	. "github.com/A-Simple-Tictok-Project/douyin/utils"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Comment struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	Content   string         `json:"title" gorm:"column:title"`
	UserId    int64          `json:"user_id" gorm:"column:user_id"`
	VideoId   int64          `json:"video_id" gorm:"column:video_id"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

type CommentDao struct{}

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

// CreateComment 评论接口
func (*CommentDao) CreateComment(comment *Comment) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		if err := tx.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteComment 删除评论
func (*CommentDao) DeleteComment(comment *Comment) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := db.Where("user_id = ? and video_id = ?", comment.UserId, comment.VideoId).Delete(comment).Error; err != nil {
			return err
		}
		if err := tx.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// QueryCommentByVideoId 评论列表
func (*CommentDao) QueryCommentByVideoId(videoId int64) ([]*Comment, error) {
	comments := make([]*Comment, 0)
	if err := db.Where("video_id = ?", videoId).Find(&comments).Error; err != nil {
		Logger.Error("query comment by video id err: " + err.Error())
		return comments, err
	}
	return comments, nil
}
