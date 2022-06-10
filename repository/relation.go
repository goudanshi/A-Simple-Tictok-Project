package repository

import "strconv"

type Relation struct {
	Id         int64 `gorm:"column:id"`
	FollowId   int64 `gorm:"column:follow_id"`
	FollowerId int64 `gorm:"column:follower_id"`
}

type RelationDao struct{}

var relationDao = RelationDao{}

func (*RelationDao) Add(relation *Relation) (int64, error) {
	err := db.Table("relation").Create(relation).Error
	if err != nil {
		return -1, err
	}
	return relation.Id, nil
}

func (*RelationDao) GetFollowIdList(followerId int64) ([]int64, error) {
	var followIds []int64
	err := db.Table("relation").Select("follow_id").Where("follower_id = ?", followerId).Find(&followIds).Error
	if err != nil {
		return nil, err
	}
	return followIds, nil
}

func (*RelationDao) GetFollowList(followerId int64) ([]User, error) {
	var follows []User
	err := db.Raw("SELECT user.id, username, password, follow_count, follower_count FROM relation INNER JOIN user on follow_id = user.id WHERE follower_id = " + strconv.FormatInt(followerId, 10)).Scan(&follows).Error
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (*RelationDao) GetFollowerIdList(followId int64) ([]int64, error) {
	var followerIds []int64
	err := db.Table("relation").Select("follower_id").Where("follow_id = ?", followId).Find(&followerIds).Error
	if err != nil {
		return nil, err
	}
	return followerIds, nil
}

func (*RelationDao) GetFollowerList(followId int64) ([]User, error) {
	var followers []User
	err := db.Raw("SELECT user.id, username, password, follow_count, follower_count FROM relation INNER JOIN user on follower_id = user.id WHERE follow_id = " + strconv.FormatInt(followId, 10)).Scan(&followers).Error
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (*RelationDao) DeleteByFollowAndFollower(followId int64, followerId int64) error {
	return db.Exec("DELETE FROM relation WHERE follow_id = " + string(followId) + " follower_id = " + string(followerId)).Error
}

func (*RelationDao) QueryByFollowIdAndFollowerId(followId int64, followerId int64) (*Relation, error) {
	relation := Relation{Id: -1}
	err := db.Table("relation").Where("follow_id = ? and follower_id = ?", followId, followerId).Find(&relation).Error
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func GetRelationDaoInstance() *RelationDao {
	return &relationDao
}
