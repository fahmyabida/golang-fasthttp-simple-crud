package main

import (
	"fasthttp_crud/handler/http"
	"fasthttp_crud/handler/mapper"
	"fasthttp_crud/handler/repository"
	"fasthttp_crud/handler/usecase"
	"fasthttp_crud/util/log"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func RunApplication() {
	fmt.Println("SIMPLE-CRUD + LOGIN SERVICE")
	loadEnv()
	serviceProperties := loadProperties()
	fmt.Println(serviceProperties)
	log.SetupLogging("./log")
	// dbClient := databaseConnect(serviceProperties)
	dbClient := &pg.DB{}

	iUserMapper := mapper.NewUserMapper()

	iUserRepo := repository.NewUserRepo(dbClient)

	iLoginUsecase := usecase.NewLoginUsecase(iUserRepo)
	iUserUsecase := usecase.NewUserUsecase(iUserMapper, iUserRepo)

	iServerHttp := http.NewServerHttp(iLoginUsecase, iUserUsecase)
	iServerHttp.Routing()
}
