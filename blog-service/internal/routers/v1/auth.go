package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/internal/service"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
	"github.com/yann0917/go-tour-book/blog-service/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	resp := app.NewResp(c)
	valid, errs := app.BindValid(c, &param)
	if !valid {
		resp.ToErrorResp(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		resp.ToErrorResp(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		fmt.Println(err)
		resp.ToErrorResp(errcode.UnauthorizedTokenGenerate)
		return
	}

	resp.ToResp(gin.H{
		"token": token,
	})
	return
}
