package service

import (
	"time"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte = []byte("xxxxxx")

type Claims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

func Register(name, password string) (uint64, bool) {
	var count int64
	config.Db.Model(&model.UserDB{}).Where("username = ?", name).Count(&count)
	if count > 0 {
		return 0, false
	}
	user := model.UserDB{
		Name:     name,
		Username: name,
		Password: password,
	}
	row := config.Db.Create(&user).RowsAffected
	if row == 0 {
		return 0, false
	}
	return user.Id, true
}

func GenerateToken(Id uint64) (string, error) {
	claims := Claims{Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 2*60*60, //过期时间 2小时
			Issuer: "Leospard", //签发人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //指定签名方法
	token, err := tokenClaims.SignedString(jwtSecret) //使用指定的secret生成token
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func Login(name string, password string) (model.User, bool) {
	var userDB model.UserDB
	var user model.User
	row := config.Db.Where("username = ? and password = ?", name, password).First(&userDB).RowsAffected
	if row == 0 {
		return model.User{}, false
	} else {
		config.Db.Model(&model.UserDB{}).Where("id = ?", userDB.Id).First(&user)
		return user, true
	}
}
