package controller

import (
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

var usersLoginInfo map[string]model.User = make(map[string]model.User)

type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// token := username + password
	uid, ok := service.Register(username, password)
	if !ok {
		println("Register fail")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Register fail"},
		})
		return
	}

	token, err := service.GenerateToken(uid)
	if err != nil {
		println("Generate token fail")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Generate token fail"},
		})
		return
	}

	newUser := model.User{
		Id:   uid,
		Name: username,
	}
	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   uid,
		Token:    token,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, ok := service.Login(username, password)
	if !ok {
		println("Login fail")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Login fail! Check your username and password"},
		})
		return
	}
	// token := username + password
	token, err := service.GenerateToken(user.Id)
	if err != nil {
		println("Generate token fail", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Generate token fail"},
		})
		return
	}
	usersLoginInfo[token] = user
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "Login Successfully"},
		UserId:   user.Id,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "No Token!"})
		return
	} else {
		claims, err := service.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Parse Token fail!"})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Token Timeout! Please login again"})
			return
		}
	}

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Please Login again"},
		})
	}
}
