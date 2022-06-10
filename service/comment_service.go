package service

import (
	"douyin/repository"
	"sync"
	"time"
)

type Comment struct {
	Id         int64     `json:"id"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
	User       User      `json:"user"`
}

func NewComment(comment *repository.Comment) (int64, error) {
	id, err := repository.GetCommentDaoInstance().Add(comment)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func DeleteComment(commentId int64) error {
	return repository.GetCommentDaoInstance().Delete(commentId)
}

func GetCommentListByVideo(videoId int64, userId int64) ([]Comment, error) {

	var data []repository.CommentUser
	var followMap map[int64]struct{}
	var dataErr, followErr error
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	data, dataErr = repository.GetCommentDaoInstance().QueryWithUser(videoId)
	waitGroup.Done()
	followMap, followErr = getFollowMap(userId)
	waitGroup.Done()

	if dataErr != nil {
		return nil, dataErr
	}
	if followErr != nil {
		return nil, followErr
	}

	result := make([]Comment, len(data))
	for i, item := range data {
		isFollow := false
		if _, ok := followMap[item.User.Id]; ok {
			isFollow = true
		}
		result[i] = Comment{
			Id:         item.Comment.Id,
			Content:    item.Comment.Content,
			CreateDate: item.Comment.CreateDate,
			User:       *convertUser(&item.User, isFollow),
		}
	}

	return result, nil
}
