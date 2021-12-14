package user

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"

	"github.com/VuHoang19/vcamAIGeneral/common"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var test_db *gorm.DB

func newUserModel() UserModel {
	return UserModel{
		ID:           2,
		Username:     "asd123!@#ASD",
		Email:        "wzt@g.cn",
		PasswordHash: "",
	}
}

func TestUserModel(t *testing.T) {
	asserts := assert.New(t)

	// Testing UserModel's password feature
	userModel := newUserModel()
	err := userModel.checkPassword("")
	asserts.Error(err, "empty password can not be set null")

	userModel = newUserModel()
	err = userModel.setPassword("")
	asserts.Error(err, "empty password can not be set null")

	userModel = newUserModel()
	err = userModel.setPassword("asd123!@ASD")
	asserts.NoError(err, "password should be set successful")
	asserts.Len(userModel.PasswordHash, 60, "password hash length should be 60")

	err = userModel.checkPassword("sd123!@#ASD")
	asserts.Error(err, "password should be checked and not validated")

	err = userModel.checkPassword("asd123!@ASD")
	asserts.NoError(err, "password should be checked and validated")
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	common.InitTestDB()
	TestAutoMigrate()
	exitVal := m.Run()
	os.Exit(exitVal)
}
