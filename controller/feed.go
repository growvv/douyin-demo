package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var latestTime int64
	latestTimeString := c.Query("latest_time")

	// 使用token去得到访问feed流的用户，去判断是否点了favorite, 没登录也可以访问。
	// Todo
	token := c.Query("token")
	user, exist := usersLoginInfo[token]
	if !exist {
		// ...
	}

	if latestTimeString != "" {
		var err error
		latestTime, err = strconv.ParseInt(latestTimeString, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Fail to parse latestTime!",
			})
			return
		}
	} else {
		latestTime = time.Now().Unix()
	}

	if videos, nextTime := service.Feed(user, latestTime); len(videos) > 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "No video availabe!",
		})
	}

}
