package service

import (
	"crypto/md5"
	"douyin/repository"
	"douyin/util"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"unsafe"
)

var authMiddleware *jwt.GinJWTMiddleware

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func GetAuthInstance() *jwt.GinJWTMiddleware {
	return authMiddleware
}

func Auth(c *gin.Context) {
	authMiddleware.MiddlewareFunc()(c)
}

func GetUserId(c *gin.Context) int64 {

	id := authMiddleware.IdentityHandler(c)

	return int64(id.(float64))
}

func GetUserInfo(id int64) (*repository.User, error) {
	return repository.GetUserDaoInstance().QueryById(id)
}

func Login(username string, password string) (*repository.User, error) {
	user, err := repository.GetUserDaoInstance().QueryByUsername(username)
	if err != nil {
		util.Logger.Error(err.Error())
		return nil, err
	}
	pass := cryptoMd5(password)

	if user.Username == username && user.Password == pass {
		return user, nil
	}
	return nil, nil
}

func Register(user *repository.User) (int64, string, error) {
	user.Password = cryptoMd5(user.Password)
	user.FollowCount = 0
	user.FollowerCount = 0

	id, err := repository.GetUserDaoInstance().Add(user)
	if err != nil {
		return -1, "", err
	}

	token, _, err := authMiddleware.TokenGenerator(id)

	if err != nil {
		return -1, "", err
	}

	return id, token, nil
}

func cryptoMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	//result := h.Sum(nil)
	//return *(*string)(unsafe.Pointer(&result))
	return hex.EncodeToString(h.Sum(nil))
}

func AuthMiddlewareInit() error {
	var err error
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         []byte("douyin-bapute"),
		TokenLookup: "query:token",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var request repository.UserRequest
			c.ShouldBind(&request)
			res, err := Login(request.Username, request.Password)
			if err != nil {
				return nil, err
			}
			if res != nil {
				return res.Id, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status_code":    1,
				"status_message": message,
			})
		},
	})
	if err != nil {
		return err
	}
	return authMiddleware.MiddlewareInit()
}

func convertUser(user *repository.User, isFollow bool) *User {
	return &User{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
}
