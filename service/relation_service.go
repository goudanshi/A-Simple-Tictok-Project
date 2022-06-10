package service

import (
	"douyin/repository"
	"sync"
)

func NewRelation(relation *repository.Relation) (int64, error) {
	//todo: 给用户的count++
	return repository.GetRelationDaoInstance().Add(relation)
}

func DeleteRelation(followId int64, followerId int64) error {
	//todo: 给count--
	return repository.GetRelationDaoInstance().DeleteByFollowAndFollower(followId, followerId)
}

func GetFollowList(userId int64) ([]User, error) {
	data, err := repository.GetRelationDaoInstance().GetFollowList(userId)
	if err != nil {
		return nil, err
	}
	result := make([]User, len(data))
	for i, item := range data {
		result[i] = *convertUser(&item, true)
	}

	return result, err
}

func getFollowMap(followerId int64) (map[int64]struct{}, error) {
	data, err := repository.GetRelationDaoInstance().GetFollowIdList(followerId)
	if err != nil {
		return nil, err
	}
	followMap := make(map[int64]struct{})
	for _, item := range data {
		followMap[item] = struct{}{}
	}
	return followMap, nil
}

func GetFollowerList(userId int64) ([]User, error) {
	dao := repository.GetRelationDaoInstance()
	var data []repository.User
	var followMap map[int64]struct{}
	var dataErr, followErr error
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go func() {
		data, dataErr = dao.GetFollowerList(userId)
		waitGroup.Done()
	}()
	go func() {
		followMap, followErr = getFollowMap(userId)
		waitGroup.Done()
	}()
	waitGroup.Wait()

	if dataErr != nil {
		return nil, dataErr
	}
	if followErr != nil {
		return nil, followErr
	}

	result := make([]User, len(data))
	for i, item := range data {
		isFollow := false
		if _, ok := followMap[item.Id]; ok {
			isFollow = true
		}
		result[i] = *convertUser(&item, isFollow)
	}

	return result, nil
}
