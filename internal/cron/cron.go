package cron

import (
	"fmt"

	"go.uber.org/zap"
)

// InitCronJobs 初始化定时任务
//
//	author centonhuang
//	update 2024-12-05 16:14:59
func InitCronJobs() {
	cleanExpiredURLsCron := NewCleanExpiredURLsCron()
	cleanExpiredURLsCron.Start()
}

type cronLoggerAdapter struct {
	prefix string
	logger *zap.Logger
}

func newCronLoggerAdapter(prefix string, logger *zap.Logger) cronLoggerAdapter {
	if prefix == "" {
		prefix = "[Cron]"
	}
	return cronLoggerAdapter{prefix: prefix, logger: logger}
}

// Error 记录错误日志
//
//	receiver l cronLoggerAdapter
//	param err error
//	param msg string
//	param keysAndValues ...interface{}
//	author centonhuang
//	update 2024-12-05 16:15:08
func (l cronLoggerAdapter) Error(err error, msg string, keysAndValues ...interface{}) {
	zapKeyValues := []zap.Field{zap.Error(err)}
	zapKeyValues = append(zapKeyValues, convertZapKeyValues(keysAndValues...)...)
	l.logger.Error(fmt.Sprintf("[%s] %s", l.prefix, msg), zapKeyValues...)
}

// Info 记录信息日志
//
//	receiver l cronLoggerAdapter
//	param msg string
//	param keysAndValues ...interface{}
//	author centonhuang
//	update 2024-12-05 16:15:25
func (l cronLoggerAdapter) Info(msg string, keysAndValues ...interface{}) {
	zapKeyValues := convertZapKeyValues(keysAndValues...)
	l.logger.Info(fmt.Sprintf("[%s] %s", l.prefix, msg), zapKeyValues...)
}

func convertZapKeyValues(keysAndValues ...interface{}) []zap.Field {
	if len(keysAndValues)%2 != 0 {
		panic("keysAndValues must be a slice of key-value pairs")
	}
	len := len(keysAndValues) / 2
	zapKeyValues := make([]zap.Field, 0, len)
	for i := 0; i < len; i++ {
		key, value := keysAndValues[i*2].(string), keysAndValues[i*2+1]
		zapKeyValues = append(zapKeyValues, zap.Any(key, value))
	}
	return zapKeyValues
}
