//+build wireinject

package router

import (
	"github.com/cwww3/im/internal/controller"
	"github.com/cwww3/im/internal/dao/db"
	"github.com/cwww3/im/internal/repository"
	"github.com/cwww3/im/internal/service"
	"github.com/google/wire"
)

func buildUserController() *controller.UserController {
	panic(wire.Build(
		controller.NewUserController,
		service.NewUserService,
		repository.NewUserRepository,
		db.GetDB,
		wire.Bind(new(service.IUserService), new(*service.UserService)),
	))
}

func buildMessageController() *controller.MessageController {
	panic(wire.Build(
		controller.NewMessageController,
		service.NewMessageService,
		repository.NewMessageRepository,
		db.GetDB,
		wire.Bind(new(service.IMessageService), new(*service.MessageService)),
	))
}
