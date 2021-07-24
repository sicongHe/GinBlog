package setting_test

import (
	"github.com/siconghe/blog/pkg/setting"
	"testing"
	"time"
)

var(
	runMode = "debug"
	hTTPPort = 5000
	readTimeout = 10 * time.Second
	writeTimeout = 10 * time.Second
	pageSize = 10
	jwtSecret = "12345$123456"
)

func TestSetting(t *testing.T) {
	t.Run("导入配置：运行模式", func(t *testing.T) {
		if runMode != setting.RunMode {
			t.Error("运行模式配置导入失败")
		}
	})
	
	t.Run("导入配置：app", func(t *testing.T) {
		if pageSize != setting.PageSize || jwtSecret != setting.JwtSecret {
			t.Error("app配置导入失败")
		}
	})

	t.Run("导入配置：server", func(t *testing.T) {
		if hTTPPort != setting.HTTPPort ||  readTimeout != setting.ReadTimeout || writeTimeout != setting.WriteTimeout{
			t.Error("server配置导入失败")
		}
	})
}
