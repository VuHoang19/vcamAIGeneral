package main

import (
	"github.com/VuHoang19/vcamAIGeneral/common"
	"github.com/VuHoang19/vcamAIGeneral/routes"
	"github.com/VuHoang19/vcamAIGeneral/user"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func Migrate() {
	user.AutoMigrate()
}

func main() {

	// load .env file from given path
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup database
	common.InitDB()
	Migrate()

	//Load the routing configuration of multiple apps
	routes.Include(user.Routes)
	//Initialize the routing of all apps
	r := routes.Init()
	if err := r.Run(); err != nil {
		log.Error("startup service failed, err:%v\n", err)
	}
}
