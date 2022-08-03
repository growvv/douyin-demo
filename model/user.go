package model

type UserDB struct {
	Id            uint64 `gorm:"column:id;autoIncrement;primaryKey,index:unique"`
	Name          string `gorm:"column:name"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   uint64 `gorm:"column:follow_count"`
	FollowerCount uint64 `gorm:"column:follower_count"`
	IsFollow      bool   `gorm:"column:is_follow"`
}

type User struct {
	Id            uint64 `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
