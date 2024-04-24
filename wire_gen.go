// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"personal_page/config"
	"personal_page/database"
	"personal_page/delivery"
	"personal_page/logger"
	"personal_page/mq"
	"personal_page/repository"
	"personal_page/usecase"
)

// Injectors from wire.go:

func InitializeServer() *delivery.WebDeli {
	viper := config.Get()
	db := database.GetDb(viper)
	userRepo := repository.NewUserRepo(db)
	userUc := usecase.NewUserUc(userRepo)
	memoRepo := repository.NewMemoRepo(db)
	memoUc := usecase.NewMemoUc(memoRepo)
	rabbitMQ := mq.GetRabbitMQ(viper)
	loggerOut := logger.NewLoggerOut(viper, rabbitMQ)
	webDeli := delivery.NewWebDeli(userUc, memoUc, loggerOut)
	return webDeli
}