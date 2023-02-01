package repository

import (
	"github.com/jason/douyin/constants"
	. "github.com/jason/douyin/utils"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Relationship struct {
	ID          int64          `json:"id" gorm:"primarykey"`
	UserId      int64          `json:"user_id" gorm:"column:user_id"`           // 女神
	SubscribeId int64          `json:"subscribe_id" gorm:"column:subscribe_id"` // 舔狗
	CreateAt    time.Time      `json:"create_at" gorm:"column:create_at"`
	UpdateAt    time.Time      `json:"update_at" gorm:"column:update_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`
}

func (r *Relationship) TableName() string {
	return constants.RelationshipTableName
}

type RelationshipDao struct{}

var (
	relationshipDao  *RelationshipDao
	relationshipOnce sync.Once
)

func NewRelationshipDaoInstance() *RelationshipDao {
	relationshipOnce.Do(func() {
		relationshipDao = &RelationshipDao{}
	})
	return relationshipDao
}

// CreateRelationship 关注接口
func (*RelationshipDao) CreateRelationship(relationship *Relationship) (int64, error) {
	// 开事务处理
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 在关系表中新建记录
		if err := tx.Create(relationship).Error; err != nil {
			return err
		}
		// 2. 为relationship.UserId用户增加一个舔狗
		if err := tx.Table("user").Where("id = ?", relationship.UserId).Update("subscribe_count", gorm.Expr("subscribe_count + ?", 1)).Error; err != nil {
			return err
		}
		// 3. 为relationship.SubscribeId用户增加一个女神
		if err := tx.Table("user").Where("id = ?", relationship.SubscribeId).Update("user_count", gorm.Expr("user_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		Logger.Error("insert relationship err: " + err.Error())
		return 0, err
	}
	return relationship.ID, nil
}

// DeleteRelationship 取关接口
func (*RelationshipDao) DeleteRelationship(relationship *Relationship) error {
	// 开事务处理
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 修改关系表中的记录
		if err := tx.Delete(relationship).Error; err != nil {
			return err
		}
		// 2. 为relationship.UserId用户减去一个舔狗
		if err := tx.Table("user").Where("id = ?", relationship.UserId).Update("subscribe_count", gorm.Expr("subscribe_count - ?", 1)).Error; err != nil {
			return err
		}
		// 3. 为relationship.SubscribeId用户减去一个女神
		if err := tx.Table("user").Where("id = ?", relationship.SubscribeId).Update("user_count", gorm.Expr("user_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// QueryRelationshipByUserId 舔狗列表
func (*RelationshipDao) QueryRelationshipByUserId(userId int64) ([]*Relationship, error) {
	relationships := make([]*Relationship, 0)
	if err := db.Where("user_id = ?", userId).Find(&relationships).Error; err != nil {
		Logger.Error("query relationship by subscribe id err: " + err.Error())
		return relationships, err
	}
	return relationships, nil
}

// QueryRelationshipBySubscribeId 女神列表
func (*RelationshipDao) QueryRelationshipBySubscribeId(subscribeId int64) ([]*Relationship, error) {
	relationships := make([]*Relationship, 0)
	if err := db.Where("subscribe_id = ?", subscribeId).Find(&relationships).Error; err != nil {
		Logger.Error("query relationship by user id err: " + err.Error())
		return relationships, err
	}
	return relationships, nil
}
