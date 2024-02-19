package api

import settings_api "server/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

// 实例化对象
var ApiGroupApp = new(ApiGroup)