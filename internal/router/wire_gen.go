// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package router

import (
	"github.com/cwww3/im/internal/controller"
	"github.com/cwww3/im/internal/dao/db"
	"github.com/cwww3/im/internal/repository"
	"github.com/cwww3/im/internal/service"
)

// Injectors from inject.go:

func buildUserController() *controller.UserController {
	gormDB := db.GetDB()
	userRepository := repository.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}

func buildMessageController() *controller.MessageController {
	gormDB := db.GetDB()
	messageRepository := repository.NewMessageRepository(gormDB)
	messageService := service.NewMessageService(messageRepository)
	messageController := controller.NewMessageController(messageService)
	return messageController
}