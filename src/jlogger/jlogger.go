package jlogger

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"

	"github.com/jackSpanrrows/go-share-library/src/common"
)

var Logger *JLogger

var RemoteIp string

func init() {
	Logger = &JLogger{
		logMode:  LOG_MODE_NONE,     // 日志模式
		version:  LOG_DEBUG_VERSION, // 版本号
		logLevel: LOG_LEVEL_DEBUG,
		logType:  LOG_TYPE_CONSOLE,
	}
}

func LoadIp(_path string) {
	resp, err := ioutil.ReadFile(_path)
	if err != nil {
		RemoteIp = "nil"
		return
	}

	RemoteIp = string(resp)
	RemoteIp = strings.Replace(RemoteIp, "\n", "", -1)
}

func CreateNone(_version string) {
	Logger = &JLogger{
		logMode:  LOG_MODE_NONE,
		version:  _version,
		logLevel: LOG_LEVEL_DEBUG,
		logType:  LOG_TYPE_CONSOLE,
	}
}

func (logger *JLogger) CheckMode(level int32) {
	if logger == nil {
		return
	}

	logger.logLevel = level
}

func (logger *JLogger) LogInfo(title string, fileFormat string, args ...interface{}) {
	if logger == nil {
		return
	}

	if logger.logLevel == LOG_LEVEL_DEBUG || logger.logLevel == LOG_LEVEL_INFO {
		MakeLineHead(logger, fileFormat, LOG_LEVEL_INFO, args...)
	}
}

func (logger *JLogger) LogWarning(title string, fileFormat string, args ...interface{}) {
	if logger == nil {
		return
	}

	if logger.logLevel >= LOG_LEVEL_WARNING {
		MakeLineHead(logger, fileFormat, LOG_LEVEL_WARNING, args...)
	}
}

func (logger *JLogger) LogErr(title string, fileFormat string, args ...interface{}) {
	if logger == nil {
		return
	}

	if logger.logLevel == LOG_LEVEL_ERR || logger.logLevel == LOG_LEVEL_DEBUG {
		MakeLineHead(logger, fileFormat, LOG_LEVEL_ERR, args...)
	}
}

func (logger *JLogger) LogDebug(title string, fileFormat string, args ...interface{}) {
	if logger == nil {
		return
	}
	if logger.logLevel == LOG_LEVEL_DEBUG {
		MakeLineHead(logger, fileFormat, LOG_LEVEL_DEBUG, args...)
	}
}

/**
  日志行头
*/
func MakeLineHead(logger *JLogger, fileFormat string, level uint8, args ...interface{}) string {
	if logger == nil {
		return ""
	}
	var logContext string
	nowTime, _ := common.GetTimeFormat(0, "01-02 15:04:05")
	_, file, line, ok := runtime.Caller(2)
	if ok {
		_, filename := path.Split(file)
		linehead := fmt.Sprintf("[%s %s:%d>>%s][%v]%s", nowTime, filename, line, logger.version, RemoteIp, fileFormat)
		logContext = fmt.Sprintf(linehead, args...)
	} else {
		logContext = fmt.Sprintf(nowTime+" >> "+logger.version+" |"+fileFormat, args...)
	}
	color := COLOR_WHITE
	switch level {
	case LOG_LEVEL_INFO:
		color = COLOR_MAGERTA
	case LOG_LEVEL_WARNING:
		color = COLOR_YELLOW
	case LOG_LEVEL_ERR:
		color = COLOR_ORINGE
	case LOG_LEVEL_DEBUG:
		color = COLOR_BLUE
	}
	fmt.Printf("\x1b[0;%dm%v\x1b[0m\n", color, logContext)

	return logContext
}
