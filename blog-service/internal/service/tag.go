package service

import (
	"github.com/yann0917/go-tour-book/blog-service/internal/model"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
)

type CountTagReq struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListReq struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagReq struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagReq struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagReq struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(params *CountTagReq) (int, error) {
	return svc.dao.CountTag(params.Name, params.State)
}

func (svc *Service) GetTagList(params *TagListReq, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(params.Name, params.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(params *CreateTagReq) error {
	return svc.dao.CreateTag(params.Name, params.State, params.CreatedBy)
}

func (svc *Service) UpdateTag(params *UpdateTagReq) error {
	return svc.dao.UpdateTag(params.ID, params.Name, params.State, params.ModifiedBy)
}

func (svc *Service) DeleteTag(params *DeleteTagReq) error {
	return svc.dao.DeleteTag(params.ID)
}
