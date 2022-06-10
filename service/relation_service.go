package service

import (
	"douyin/repository"
	"douyin/util"
	"sync"
)

func NewRelation(relation *repository.Relation) (int64, error) {
	id, err := repository.GetRelationDaoInstance().Add(relation)
	if err != nil {
		return -1, err
	}
	go func(followId int64) {
		userDao := repository.GetUserDaoInstance()
		user, e := userDao.QueryById(followId)
		if e != nil {
			util.Logger.Error(e.Error())
		}
		user.FollowCount++
		e = userDao.Update(user)
		if e != nil {
			return
		}
	}(relation.FollowId)
	go func(followerId int64) {
		userDao := repository.GetUserDaoInstance()
		user, e := userDao.QueryById(followerId)
		if e != nil {
			util.Logger.Error(e.Error())
		}
		user.FollowerCount++
		e = userDao.Update(user)
		if e != nil {
			return
		}
	}(relation.FollowerId)
	return id, nil
}

func DeleteRelation(followId int64, followerId int64) error {
	err := repository.GetRelationDaoInstance().DeleteByFollowAndFollower(followId, followerId)
	if err != nil {
		return err
	}
	go func(followId int64) {
		userDao := repository.GetUserDaoInstance()
		user, e := userDao.QueryById(followId)
		if e != nil {
			util.Logger.Error(e.Error())
		}
		user.FollowCount--
		e = userDao.Update(user)
		if e != nil {
			return
		}
	}(followId)
	go func(followerId int64) {
		userDao := repository.GetUserDaoInstance()
		user, e := userDao.QueryById(followerId)
		if e != nil {
			util.Logger.Error(e.Error())
		}
		user.FollowerCount--
		e = userDao.Update(user)
		if e != nil {
			return
		}
	}(followerId)
	return nil
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
