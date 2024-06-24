package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"volcano.user_srv/global"
)

var DefaultGormLogger = MysqlGormLogger{
	// 安静模式，这里保证 gorm 内部不会打印，我们已经实现了他的全部方法
	LogLevel:logger.Silent,
	SlowThreshold:200 * time.Millisecond,
}

type MysqlGormLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func (mgl *MysqlGormLogger) LogMode(logLevel logger.LogLevel) logger.Interface {
	mgl.LogLevel = logLevel
	return mgl
}

func (mgl *MysqlGormLogger) Info(ctx context.Context, message string, values ...interface{}) {
	zap.S().Infow("_com_mysql_Info", "message", message, "values", fmt.Sprint(values...))
}

func (mgl *MysqlGormLogger) Warn(ctx context.Context, message string, values ...interface{}) {
	zap.S().Infow("_com_mysql_Warn", "message", message, "values", fmt.Sprint(values...))
}

func (mgl *MysqlGormLogger) Error(ctx context.Context, message string, values ...interface{}) {
	zap.S().Infow("_com_mysql_Error", "message", message, "values", fmt.Sprint(values...))
}

// .Debug() 时调用
func (mgl *MysqlGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {

	if mgl.LogLevel <= logger.Silent {
		return
	}

	sqlStr, rows := fc()
	currentTime := begin.Format(global.TimeFormat)
	elapsed := time.Since(begin)
	msg := map[string]interface{}{
		"FileWithLineNum": utils.FileWithLineNum(),
		"sql":             sqlStr,
		"rows":            "-",
		"proc_time":       float64(elapsed.Milliseconds()),
		"current_time":    currentTime,
	}
	switch {
	case err != nil && mgl.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		msg["err"] = err
		if rows == -1 {
			zap.S().Errorw("_com_mysql_failure", "message", msg)
		} else {
			msg["rows"] = rows
			zap.S().Infow("_com_mysql_failure", "message", msg)
		}
	case elapsed > mgl.SlowThreshold && mgl.SlowThreshold != 0 && mgl.LogLevel >= logger.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", mgl.SlowThreshold)
		msg["slowLog"] = slowLog
		if rows == -1 {
			zap.S().Infow("_com_mysql_failure", "message", msg)
		} else {
			msg["rows"] = rows
			zap.S().Infow("_com_mysql_failure", "message", msg)
		}
	case mgl.LogLevel == logger.Info:
		if rows == -1 {
			zap.S().Infow("_com_mysql_failure", "message", msg)
		} else {
			msg["rows"] = rows
			zap.S().Infow("_com_mysql_failure", "message", msg)
		}
	}
}
