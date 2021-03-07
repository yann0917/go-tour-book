package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/yann0917/go-tour-book/blog-service/docs"
	v1 "github.com/yann0917/go-tour-book/blog-service/internal/routers/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		tag := v1.NewTag()
		tags := apiv1.Group("/tags")
		{
			tags.POST("", tag.Create)
			tags.DELETE("/:id", tag.Delete)
			tags.PUT("/:id", tag.Update)
			tags.PATCH("/:id/state", tag.Update)
			tags.GET("/:id", tag.Get)
			tags.GET("", tag.List)
		}
		article := v1.NewArticle()
		articles := apiv1.Group("/articles")
		{
			articles.POST("", article.Create)
			articles.DELETE("/:id", article.Delete)
			articles.PUT("/:id", article.Update)
			articles.PATCH("/:id/state", article.Update)
			articles.GET("/:id", article.Get)
			articles.GET("", article.List)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
