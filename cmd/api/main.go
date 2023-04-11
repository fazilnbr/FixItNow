package main

import (
	"fmt"
	"log"

	_ "github.com/fazilnbr/project-workey/cmd/api/docs"
	_ "github.com/fazilnbr/project-workey/pkg/domain"
	_ "github.com/fazilnbr/project-workey/pkg/utils"

	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/db"

	"github.com/fazilnbr/project-workey/pkg/di"
)

// @title Go + Gin Workey API
// @version 1.0
// @description This is a sample server Job Portal server. You can visit the GitHub repository at https://github.com/fazilnbr/Job_Portal_Project

// @contact.name API Support
// @contact.url https://fazilnbr.github.io/mypeosolal.web.portfolio/
// @contact.email fazilkp2000@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	fmt.Println("starting my job-portal project in clean code architecture")

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	// db.ConnectDB(config)
	gorm, _ := db.ConnectGormDB(config)
	fmt.Printf("\ngorm : %v\n\n", gorm)

	server, diErr := di.InitializeAPI(config)
	fmt.Printf("\n\n\nserver ; %v\n\n\n", server)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
