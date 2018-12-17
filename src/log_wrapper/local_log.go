package log_wrapper

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"time"
	"path"
	"io/ioutil"
	
	"io"
	"github.com/sirupsen/logrus"
)

/******************************************************************************
 * create new local file log_wrapper
 ******************************************************************************/
func NewLocalLogger(logPath string, logName string, logLevel logrus.Level, maxAge time.Duration,
	rotationTime time.Duration) (error, *logrus.Logger) {
	
	// can not access dir:logPath
	_, err := ioutil.ReadDir(logPath)
	if err != nil {
		return err, nil
	}
	
	baseName := path.Join(logPath, logName)
	
	writer, err := rotatelogs.New(
		baseName+"..%Y%m%d%H%M%S",
		rotatelogs.WithMaxAge(maxAge),             // 日志保存时间
		rotatelogs.WithRotationTime(rotationTime), // 切割时间间隔
	)
	if err != nil {
		return err, nil
	}
	
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
	
	newLog := logrus.New()
	newLog.SetOutput(writer)
	newLog.SetLevel(logLevel)
	newLog.AddHook(lfHook)
	
	return nil, newLog
}

/******************************************************************************
 * create io.Writer log_wrapper
 ******************************************************************************/
func NewWriterLogger(writer io.Writer, logLevel logrus.Level) *logrus.Logger {
	
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
	
	newLog := logrus.New()
	newLog.SetOutput(writer)
	newLog.SetLevel(logLevel)
	newLog.AddHook(lfHook)
	
	return newLog
}
