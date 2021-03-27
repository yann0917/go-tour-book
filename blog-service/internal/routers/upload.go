package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/internal/service"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
	"github.com/yann0917/go-tour-book/blog-service/pkg/convert"
	"github.com/yann0917/go-tour-book/blog-service/pkg/errcode"
	"github.com/yann0917/go-tour-book/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	resp := app.NewResp(c)

	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		errResp := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToErrorResp(errResp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		resp.ToErrorResp(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())

	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		//global.Logger.Errorf("svc.UploadFile err: %v", err)
		errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		resp.ToErrorResp(errResp)
		return
	}
	resp.ToResp(gin.H{
		"access_url": fileInfo.AccessUrl,
	})
}
