package user

import (
	"github.com/VuHoang19/vcamAIGeneral/common"
	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}
