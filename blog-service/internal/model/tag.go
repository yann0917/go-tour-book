package model

import (
	"github.com/jinzhu/gorm"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t *Tag) Count(db *gorm.DB) (count int, err error) {
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	err = db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (t *Tag) List(db *gorm.DB, pageOffset, pageSize int) (tags []*Tag, err error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return
	}
	return
}

func (t *Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t *Tag) Update(db *gorm.DB, value interface{}) error {
	db = db.Model(t).Where("id = ? and is_del = ?", t.ID, 0)
	return db.Updates(value).Error
}

func (t *Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.ID, 0).Delete(&t).Error
}
