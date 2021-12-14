package user

import (
	"github.com/VuHoang19/vcamAIGeneral/common"
	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (userModel *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, userModel)
	if err != nil {
		return err
	}
	userModel.userModel.Username = userModel.User.Username
	userModel.userModel.Email = userModel.User.Email

	if userModel.User.Password != common.NBRandomPassword {
		userModel.userModel.setPassword(userModel.User.Password)
	}

	return nil
}

func NewUserModelValidator() UserModelValidator {
	userModelvalidator := UserModelValidator{}
	//userModelValidator.User.Email ="abc@gmail.com"
	return userModelvalidator
}
