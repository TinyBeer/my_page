// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"personal_page/config"
	"personal_page/database"
	"personal_page/delivery"
	"personal_page/repository"
	"personal_page/usecase"
)

// Injectors from wire.go:

func InitializeServer() *delivery.WebDeli {
	viper := config.Get()
	db := database.GetDb(viper)
	mongoDatabase := database.GetMongoDB(viper)
	userRepository := repository.GetUserRepository(db, mongoDatabase)
	userUc := usecase.NewUserUc(userRepository)
	memoRepository := repository.GetMemoRepository(db, mongoDatabase)
	memoUc := usecase.NewMemoUc(memoRepository)
	typedClient := database.GetES(viper)
	movieRepository := repository.GetMovieRepository(typedClient)
	movieUc := usecase.NewMovieUc(movieRepository)
	webDeli := delivery.NewWebDeli(viper, userUc, memoUc, movieUc)
	return webDeli
}
