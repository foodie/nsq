// short for "log"
package lg

import (
	"fmt"
	"strings"
)

//定义LogLevel的类型
type LogLevel int

const (
	DEBUG = LogLevel(1)
	INFO  = LogLevel(2)
	WARN  = LogLevel(3)
	ERROR = LogLevel(4)
	FATAL = LogLevel(5)
)

type AppLogFunc func(lvl LogLevel, f string, args ...interface{})

//Logger接口实现了Output类
type Logger interface {
	Output(maxdepth int, s string) error
}

//空的Logger
type NilLogger struct{}

func (l NilLogger) Output(maxdepth int, s string) error {
	return nil
}

//LogLevel转换成对应的数字
func (l LogLevel) String() string {
	switch l {
	case 1:
		return "DEBUG"
	case 2:
		return "INFO"
	case 3:
		return "WARNING"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	}
	panic("invalid LogLevel")
}

//通过string获取对应的LogLevel
func ParseLogLevel(levelstr string, verbose bool) (LogLevel, error) {
	lvl := INFO

	switch strings.ToLower(levelstr) {
	case "debug":
		lvl = DEBUG
	case "info":
		lvl = INFO
	case "warn":
		lvl = WARN
	case "error":
		lvl = ERROR
	case "fatal":
		lvl = FATAL
	default:
		return lvl, fmt.Errorf("invalid log-level '%s'", levelstr)
	}
	if verbose {
		lvl = DEBUG
	}
	return lvl, nil
}

//通过Logger打印对应基本的日志，cfg配置级别，msgLevel消息级别
func Logf(logger Logger, cfgLevel LogLevel, msgLevel LogLevel, f string, args ...interface{}) {
	if cfgLevel > msgLevel {
		return
	}
	logger.Output(3, fmt.Sprintf(msgLevel.String()+": "+f, args...))
}
