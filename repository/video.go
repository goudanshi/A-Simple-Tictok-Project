package repository

import (
	"strconv"
	"time"
)

type Video struct {
	Id            int64     `json:"id" gorm:"column:id"`
	PublisherId   int64     `json:"publisher_id" gorm:"column:publisher_id"`
	Title         string    `json:"title" gorm:"column:title"`
	VideoUrl      string    `json:"play_url" gorm:"column:video_url"`
	CoverUrl      string    `json:"cover_url" gorm:"column:cover_url"`
	FavoriteCount int64     `json:"favorite_count" gorm:"column:favorite_count"`
	CommentCount  int64     `json:"comment_count" gorm:"column:comment_count"`
	CreateDate    time.Time `json:"create_date" gorm:"column:create_date"`
}

type VideoUser struct {
	Video Video `gorm:"embedded"`
	User  User  `gorm:"embedded"`
}

type VideoDao struct{}

var videoDao = VideoDao{}

func (*VideoDao) Add(video *Video) (int64, error) {
	err := db.Table("video").Create(video).Error
	if err != nil {
		return -1, err
	}
	return video.Id, nil
}

func (*VideoDao) QueryById(id int64) (*Video, error) {
	var video Video
	err := db.Table("video").Where("id = ?", id).Find(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (*VideoDao) QueryByPublisher(publisher int64) ([]Video, error) {
	var videos []Video
	err := db.Table("video").Where("publisher_id = ?", publisher).Order("create_date desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) Query(limit int) ([]VideoUser, error) {
	var videos []VideoUser
	err := db.Raw("SELECT video.id as id, title, video_url, cover_url, favorite_count, comment_count, create_date, user.id, username, follow_count, follower_count FROM video left join user on video.publisher_id = user.id order by create_date desc limit " + strconv.FormatInt(int64(limit), 10)).Scan(&videos).Error
	//err := db.Table("video").Select("video.id as id, video_url, cover_url, favorite_count, comment_count, title, user.id as publisher_id, user.username as username, user.follow_count as follow_count, user.follower_count as follower_count").Joins("left join user on video.publisher_id = user.id").Scan(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) QueryUserVideo(limit int64) ([]VideoUser, error) {
	var videos []VideoUser
	err := db.Raw("SELECT video.id, title, video_url, cover_url, favorite_count, comment_count, user.id, username, follow_count, follower_count FROM video INNER JOIN user ON publisher_id = user.id  order by create_date desc limit " + strconv.FormatInt(limit, 10)).Scan(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) Update(video *Video) error {
	return db.Table("video").Save(video).Error
}

func GetVideoDaoInstance() *VideoDao {
	return &videoDao
}
