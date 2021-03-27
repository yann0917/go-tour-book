package dao

import (
	"github.com/yann0917/go-tour-book/blog-service/internal/model"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(ID uint32, name string, state uint8, modifiedBy string) error {
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	tag := model.Tag{
		Model: &model.Model{
			ID: ID,
		},
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(ID uint32) error {
	tag := model.Tag{Model: &model.Model{ID: ID}}
	return tag.Delete(d.engine)
}
