package util

import "GoStartCode/data-access/GormTest/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
