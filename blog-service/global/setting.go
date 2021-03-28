package global

import (
	"github.com/yann0917/go-tour-book/blog-service/pkg/logger"
	"github.com/yann0917/go-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServiceSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)
