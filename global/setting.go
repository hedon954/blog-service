package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

/**
全局配置
*/

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
