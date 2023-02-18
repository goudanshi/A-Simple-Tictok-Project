package repository

import (
	"strconv"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	Content    string    `gorm:"column:content"`
	CreateDate time.Time `gorm:"create_date"`
}

type CommentUser struct {
	Comment Comment `gorm:"embedded"`
	User    User    `gorm:"embedded"`
}

type CommentDao struct{}

var commentDao = CommentDao{}

func (*CommentDao) Add(comment *Comment) (int64, error) {
	err := db.Table("comment").Create(comment).Error
	if err != nil {
		return -1, nil
	}
	return comment.Id, nil
}

func (*CommentDao) QueryById(id int64) (*Comment, error) {
	var comment Comment
	err := db.Table("comment").Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (*CommentDao) QueryByVideo(videoId int64) ([]Comment, error) {
	var comments []Comment
	err := db.Table("comment").Where("video_id = ?", videoId).Order("create_date desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (*CommentDao) Delete(commentId int64) error {
	return db.Table("comment").Delete(&Comment{}, commentId).Error
}

func (*CommentDao) QueryWithUser(videoId int64) ([]CommentUser, error) {
	var result []CommentUser
	err := db.Raw("SELECT * FROM comment LEFT JOIN user on comment.user_id = user.id WHERE video_id = " + strconv.FormatInt(videoId, 10)).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetCommentDaoInstance() *CommentDao {
	return &commentDao
}
