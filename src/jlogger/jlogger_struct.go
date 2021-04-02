package jlogger

import (
	"net"
	"os"
)

const (
	LOG_LEVEL_INFO    = 0
	LOG_LEVEL_WARNING = 1
	LOG_LEVEL_ERR     = 2
	LOG_LEVEL_DEBUG   = 3
)

const (
	COLOR_ORINGE = uint8(iota + 91) // 使用自增常量
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_PURPLE
	COLOR_MAGERTA
	COLOR_BLUE
	COLOR_WHITE
	COLOR_BLACK
)

const (
	LOG_TYPE_CONSOLE = 1 // 命令行输出
	LOG_TYPE_ES      = 2 // elasticsearch
	LOG_TYPE_FILE    = 4
	LOG_TYPE_ALL     = LOG_TYPE_CONSOLE | LOG_TYPE_ES | LOG_TYPE_FILE
)

const (
	LOG_MODE_FILE     = 1 // 文件
	LOG_MODE_NONE     = 2
	LOG_MODE_GRAYLOG  = 4
	LOG_MODE_JLOGGER  = 8
	LOG_DEBUG_VERSION = "v1.0.0"
)

type GrayLog struct {
	conn     net.Conn
	facility string
}

type JLogger struct {
	fileName string
	file     *os.File
	grayLog  *GrayLog
	logMode  int32
	logLevel int32
	version  string
	logType  uint32
	sip      string
	saveMode uint32
}
