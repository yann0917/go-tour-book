package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/global/errcode"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a *Article) Get(c *gin.Context) {
	app.NewResp(c).ToErrorResp(errcode.ServerError)
}
func (a *Article) List(c *gin.Context)   {}
func (a *Article) Create(c *gin.Context) {}
func (a *Article) Update(c *gin.Context) {}
func (a *Article) Delete(c *gin.Context) {}
