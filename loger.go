package logr

import (
	"runtime"
	"path/filepath"
	"fmt"
	"github.com/sirupsen/logrus"
	"bytes"
	"strconv"
)

const (
	GLOBAL = "global" //全局name空间
)

func getEntry(name string) *logrus.Entry {
	return logger.WithFields(locate(logrus.Fields{"N": name}))
}
func locate(fields logrus.Fields) logrus.Fields {
	_, path, line, ok := runtime.Caller(3)
	if ok {
		_, file := filepath.Split(path)
		fields["F"] = fmt.Sprintf("[%s:%d]", file, line)
	}
	fields["G"] = fmt.Sprintf("%d", GetGID())
	return fields
}

func Debug(msg ...interface{}) {
	getEntry(GLOBAL).Debug(msg...)
}
func Debugln(msg ...interface{}) {
	logger.Debugln(msg...)
}
func Debugf(format string, msg ...interface{}) {
	logger.Debugf(format, msg...)
}

func Info(msg ...interface{}) {
	logger.Info(msg...)
}
func Infoln(msg ...interface{}) {
	logger.Infoln(msg...)
}
func Infof(format string, msg ...interface{}) {
	logger.Infof(format, msg...)
}

func Warn(msg ...interface{}) {
	logger.Warn(msg...)
}
func Warnln(msg ...interface{}) {
	logger.Warnln(msg...)
}
func Warnf(format string, msg ...interface{}) {
	logger.Warnf(format, msg...)
}
func Error(msg ...interface{}) {
	logger.Error(msg...)
}
func Errorln(msg ...interface{}) {
	logger.Errorln(msg...)
}
func Errorf(format string, msg ...interface{}) {
	logger.Errorf(format, msg...)
}
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}