package middleware

import (
	"fmt"
	"time"

	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
	"github.com/yann0917/go-tour-book/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/global"
	"github.com/yann0917/go-tour-book/blog-service/pkg/email"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
		IsSSL:    global.EmailSetting.IsSSL,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Error(c, s, err)

				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResp(c).ToErrorResp(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}

}
