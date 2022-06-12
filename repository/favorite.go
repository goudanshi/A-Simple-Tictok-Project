package repository

import (
	"fmt"
	"strconv"
	"time"
)

type Favorite struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	CreateDate time.Time `gorm:"column:create_date"`
}

type FavoriteDao struct{}

var favoriteDao = FavoriteDao{}

func (*FavoriteDao) Add(favorite *Favorite) (int64, error) {
	err := db.Table("favorite").Create(favorite).Error
	if err != nil {
		return -1, err
	}
	return favorite.Id, nil
}

func (*FavoriteDao) QueryVideoByUserAndEarlyDate(userId int64, date time.Time) ([]int64, error) {
	var ids []int64
	err := db.Table("favorite").Select("video_id").Where("user_id = ?", userId).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (*FavoriteDao) DeleteByUserVideo(userId int64, videoId int64) error {
	sql := "DELETE FROM favorite WHERE user_id = " + strconv.FormatInt(userId, 10) + " AND video_id = " + strconv.FormatInt(videoId, 10)
	fmt.Println(sql)
	return db.Exec(sql).Error
}

func (*FavoriteDao) QueryFavoriteVideo(userId int64) ([]VideoUser, error) {
	var data []VideoUser
	err := db.Raw("SELECT video.id, title, video_url, cover_url, favorite_count, comment_count, user.id, username, follow_count, follower_count FROM favorite INNER JOIN video ON favorite.video_id = video.id INNER JOIN user ON video.publisher_id = user.id WHERE favorite.user_id = " + strconv.FormatInt(userId, 10)).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetFavoriteDaoInstance() *FavoriteDao {
	return &favoriteDao
}
