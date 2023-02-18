package repository

type User struct {
	Id            int64  `gorm:"column:id"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

type UserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserDao struct{}

var userDao = UserDao{}

func (*UserDao) Add(user *User) (int64, error) {
	err := db.Table("user").Create(user).Error
	if err != nil {
		return -1, err
	}
	return user.Id, nil
}

func (*UserDao) Update(user *User) error {
	return db.Table("user").Updates(user).Error
}

func (*UserDao) QueryById(id int64) (*User, error) {
	var user User
	err := db.Table("user").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryByUsername(username string) (*User, error) {
	var user User
	err := db.Table("user").Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserDaoInstance() *UserDao {
	return &userDao
}
