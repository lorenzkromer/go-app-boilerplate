package main

import (
	"fmt"
	"strconv"

	"github.com/lorenzkromer/go-app-boilerplate/pkg/adapters"
	"github.com/lorenzkromer/go-app-boilerplate/pkg/config"
	"github.com/lorenzkromer/go-app-boilerplate/pkg/domain"
	"github.com/lorenzkromer/go-app-boilerplate/pkg/services"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Warnf("Recovered in %s", r)
		}
	}()

	// Load configs
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Setup DB client
	dsn := "host=" + config.Config.Database.Host + " port=" + strconv.Itoa(config.Config.Database.Port) + " user=" + config.Config.Database.User + " dbname=" + config.Config.Database.Name + " password=" + config.Config.Database.Password + " sslmode=" + config.Config.Database.SslMode
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.Example{})

	// Setup repository
	exampleRepo := adapters.NewExampleRepository()

	// Setup service
	exampleService := services.NewExampleService(exampleRepo)

	// Run
	fmt.Println(exampleService.Hello())
}
