package controller

import (
	"github.com/cwww3/im/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SESSION_KEY = "user"

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Login(ctx *gin.Context) {
	res := gin.H{}

	var params = new(struct {
		Email    string
		Password string
	})
	err := ctx.Bind(params)
	if err != nil {
		res["errmsg"] = err.Error()
		ctx.HTML(http.StatusOK, "login.index", res)
	}

	user, err := uc.userService.Login(params.Email, params.Password)
	if err != nil {
		res["errmsg"] = err.Error()
		ctx.HTML(http.StatusOK, "login.index", res)
		return
	}
	session := sessions.Default(ctx)
	session.Set(SESSION_KEY, user)
	res["loginUser"] = user

	otherUsers, err := uc.userService.GetUserListExcept(user)
	if err != nil {
		res["errmsg"] = err.Error()
		ctx.HTML(http.StatusOK, "login.index", res)
		return
	}
	res["otherUsers"] = otherUsers

	contact, err := uc.userService.GetContacts(user)
	if err != nil {
		res["errmsg"] = err.Error()
		ctx.HTML(http.StatusOK, "login.index", res)
		return
	}
	res["contactVO"] = contact

	ctx.HTML(http.StatusOK, "index.html", res)
}

func (uc *UserController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(SESSION_KEY)
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func (uc *UserController) WelcomePage(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(SESSION_KEY)
	if user != nil {
		ctx.HTML(http.StatusOK, "index.html", user)
		return
	}
	ctx.HTML(http.StatusOK, "login.html", nil)
}
