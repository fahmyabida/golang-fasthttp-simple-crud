package main

import (
	"context"
	"fasthttp_crud/model"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"os"
)

func loadEnv() {
	fmt.Println("Load env variable...")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success load env variable!")
}

func loadProperties() model.ServiceProperties {
	return model.ServiceProperties{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

func databaseConnect(serviceProperties model.ServiceProperties) (dbClient *pg.DB) {
	fmt.Println("Connecting to database...")
	dbAddr := serviceProperties.DbHost
	if serviceProperties.DbPort != "" {
		dbAddr += ":" + serviceProperties.DbPort
	}
	dbClient = pg.Connect(&pg.Options{
		Addr:     dbAddr,
		User:     serviceProperties.DbUser,
		Password: serviceProperties.DbPassword,
		Database: serviceProperties.DbName,
	})
	err := dbClient.Ping(context.Background())
	if err != nil {
		fmt.Println("Failed connect to database!")
		panic(err)
	}
	fmt.Println("Connected to database!")
	return dbClient
}
