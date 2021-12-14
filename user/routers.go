package user

import (
	"net/http"

	"github.com/VuHoang19/vcamAIGeneral/common"
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	// e . GET ( "/resume" , resumeHander )
	// e . GET ( "/info" , infoHander )
	UsersRegister(e)
}

func UsersRegister(e *gin.Engine) {
	e.POST("/", UsersRegistration)
	e.POST("/login", UsersLogin)
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})

}

func UsersLogin(c *gin.Context) {

}
