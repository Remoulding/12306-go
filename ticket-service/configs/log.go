package configs

import "github.com/sirupsen/logrus"

// Log 全局日志实例
var Log = logrus.New()

func InitLog() {
	Log.SetReportCaller(true) // 启用调用者报告
}
