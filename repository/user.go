package repository

import (
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	. "github.com/A-Simple-Tictok-Project/douyin/utils"
	"gorm.io/gorm"
	"sync"
	"time"
)

type User struct {
	ID             int64          `json:"id" gorm:"primarykey"`
	Username       string         `json:"username" gorm:"column:username"`
	Password       string         `json:"password" gorm:"column:password"`
	UserCount      int64          `json:"user_count" gorm:"column:user_count;default:0"`           // 女神数
	SubscribeCount int64          `json:"subscribe_count" gorm:"column:subscribe_count;default:0"` // 舔狗数
	CreatedAt      time.Time      `json:"created_at" gorm:"column:create_at"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeleteAt       gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`

	// 外键字段
	Subscribes []*Relationship `json:"subscribes" gorm:"foreignKey:SubscribeId;references:ID"` // 我的舔狗
	Users      []*Relationship `json:"fans" gorm:"foreignKey:UserId;references:ID"`            // 我的女神
	Videos     []*Video        `json:"videos" gorm:"foreignKey:UserId;references:ID"`
	Comments   []*Comment      `json:"comments" gorm:"foreignKey:UserId;references:ID"`
	Likes      []*Like         `json:"likes" gorm:"foreignKey:UserId;references:ID"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

type UserDao struct{}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// MQueryUserById 用户信息
func (*UserDao) MQueryUserById(ids []int64) ([]*User, error) {
	var users []*User
	err := db.Preload("Subscribes").Preload("Users").Preload("Videos").Preload("Comments").Preload("Likes").Where("id in (?)", ids).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return users, nil
	}
	if err != nil {
		Logger.Error("query users by id err: " + err.Error())
		return users, err
	}
	return users, nil
}

// QueryUserByUsername 登录接口
func (*UserDao) QueryUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		Logger.Error("find user by username err: " + err.Error())
		return nil, err
	}
	return &user, nil
}

// CreateUser 注册接口
func (d *UserDao) CreateUser(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		Logger.Error("insert user err: " + err.Error())
		return err
	}
	return nil
}
