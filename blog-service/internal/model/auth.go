package model

import "github.com/jinzhu/gorm"

type Auth struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	*Model
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (auth Auth, err error) {
	err = db.
		Where("app_key = ? and app_secret = ? and is_del = ?", a.AppKey, a.AppSecret, 0).
		First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}
