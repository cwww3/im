package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("userSession", store))

	userCtl := buildUserController()
	// 用户
	{
		userGroup := e.Group("")
		userGroup.GET("/", userCtl.WelcomePage)
		userGroup.POST("/login", userCtl.Login)
		userGroup.GET("/logout", userCtl.Logout)
	}

	// 消息
	messageCtl := buildMessageController()
	{
		msgGroup := e.Group("")
		msgGroup.POST("/sendMessage", messageCtl.SendMessage)
		msgGroup.GET("/queryMsg", messageCtl.QueryMsg)
		msgGroup.GET("/queryMsgSinceMid", messageCtl.QueryMsgSinceMid)
		msgGroup.GET("/queryContacts", messageCtl.QueryContacts)
	}

}
