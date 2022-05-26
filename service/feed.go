package service

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
)

func Feed(user model.User, latestTime int64) ([]model.Video, int64) {
	var videosDB []model.VideoDB
	var videos []model.Video
	config.Db.Model(&model.VideoDB{}).Where("create_time < ?", latestTime).Order("create_time DESC").Limit(30).Find(&videosDB)
	var nextTime int64
	if len(videosDB) > 0 {
		nextTime = videosDB[len(videosDB)-1].CreateTime
	}
	for _, video := range videosDB {
		var author model.User
		// 通过AuthorId查找视频作者相关信息
		config.Db.Model(&model.UserDB{}).Where("id = ?", video.AuthorId).First(&author)
		// Todo 通过user查出视频是否点了赞

		videos = append(videos, model.Video{
			Id:            video.Id,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			// Todo
			IsFavorite: false,
		})
	}
	return videos, nextTime
}
