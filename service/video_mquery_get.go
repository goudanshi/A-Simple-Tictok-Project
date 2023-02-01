package service

import (
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
	"strconv"
	"time"
)

type VideoMQueryGetFlow struct {
	// 请求需要的数据
	LatestTime string `json:"latest_time"`
	// 响应需要的数据
	Videos    []*repository.Video        `json:"videos"`
	AuthorMap map[int64]*repository.User `json:"author_map"`
}

func NewVideoMQueryGetFlow(latestTime string) *VideoMQueryGetFlow {
	return &VideoMQueryGetFlow{LatestTime: latestTime}
}

func VideoMQueryGet(latestTime string) (*VideoMQueryGetFlow, errno.ErrNo) {
	return NewVideoMQueryGetFlow(latestTime).Do()
}

func (f *VideoMQueryGetFlow) Do() (*VideoMQueryGetFlow, errno.ErrNo) {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. query video information
	if err := f.queryVideo(); err.ErrCode != 0 {
		return nil, err
	}
	// 3. query user information
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	return f, errno.Success
}
func (f *VideoMQueryGetFlow) checkParam() errno.ErrNo {
	if f.LatestTime == "" {
		f.LatestTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		// latestTime是时间戳
		latestTimeStr, err := strconv.ParseInt(f.LatestTime, 10, 64)
		if err != nil {
			return errno.ParamErr
		}
		f.LatestTime = time.UnixMilli(latestTimeStr).Format("2006-01-02 15:04:05")
	}
	return errno.Success
}

func (f *VideoMQueryGetFlow) queryVideo() errno.ErrNo {
	videos, err := repository.NewVideoDaoInstance().QueryVideoByTime(f.LatestTime)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Videos = videos
	return errno.Success
}
func (f *VideoMQueryGetFlow) queryUser() errno.ErrNo {
	userIds := make([]int64, 0)
	for _, video := range f.Videos {
		userIds = append(userIds, video.UserId)
	}
	users, err := repository.NewUserDaoInstance().MQueryUserById(userIds)
	if err != nil {
		return errno.ConvertErr(err)
	}
	authorMap := make(map[int64]*repository.User)
	for _, user := range users {
		authorMap[user.ID] = user
	}
	f.AuthorMap = authorMap
	return errno.Success
}
