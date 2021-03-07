package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yann0917/go-tour-book/blog-service/global"
	"github.com/yann0917/go-tour-book/blog-service/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
}

func NewDBEngine(dbsetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbsetting.UserName,
		dbsetting.Password,
		dbsetting.Host,
		dbsetting.DBName,
		dbsetting.Charset,
		dbsetting.ParseTime,
	)
	db, err := gorm.Open(dbsetting.DBType, dsn)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dbsetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbsetting.MaxOpenconns)

	return db, nil

}
