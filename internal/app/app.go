package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code string `json:"code"`
	ErrMsg string `json:"errmsg"`
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context,data interface{}) {
	ctx.JSON(http.StatusOK,Response{
		Code:   "0",
		ErrMsg: "",
		Data:   data,
	})
}

func Failed(ctx *gin.Context,err error) {
	ctx.JSON(http.StatusOK,Response{
		Code:   "1",
		ErrMsg: err.Error(),
	})
}