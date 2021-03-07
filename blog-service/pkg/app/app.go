package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/global/errcode"
)

type Resp struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResp(ctx *gin.Context) *Resp {
	return &Resp{
		Ctx: ctx,
	}
}

func (r *Resp) ToResp(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Resp) ToRespList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Resp) ToErrorResp(err *errcode.Error) {
	resp := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		resp["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), resp)
}
