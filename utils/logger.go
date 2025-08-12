package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// setupLogger sets up the logger with a specific format and output file.
func SetupLogger() {
	// 设置日志输出目标为标准输出
	logrus.SetOutput(os.Stdout)
	// 设置日志级别为Info
	logrus.SetLevel(logrus.InfoLevel)

	// 创建一个TextFormatter实例
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true, // 启用彩色输出
	}
	// 设置日志格式化器
	logrus.SetFormatter(formatter)
}
