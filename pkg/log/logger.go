package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

var Logger *logger

type logger struct {
	OnFile bool
	error  *logrus.Logger
	panic  *logrus.Logger
	info   *logrus.Logger
	api    *logrus.Logger
}

func InitLogger(onFile bool) (*logger, error) {
	dir, _ := os.Getwd()
	var apiLogger *logrus.Logger
	var errLogger *logrus.Logger
	var infoLogger *logrus.Logger
	var panicLogger *logrus.Logger
	if onFile {
		//日志记录文件
		errFile := dir + "/log/error/error.log"
		panicFile := dir + "/log/panic/panic.log"
		infoFile := dir + "/log/info/info.log"
		errSrc, err := os.OpenFile(errFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}
		panicSrc, err := os.OpenFile(panicFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}
		infoSrc, err := os.OpenFile(infoFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}
		//错误等级日志
		errLogger = logrus.New()
		errLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		errLogger.SetOutput(errSrc)

		//异常等级日志
		panicLogger = logrus.New()
		panicLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		panicLogger.SetOutput(panicSrc)

		//操作等级日志
		infoLogger = logrus.New()
		infoLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		infoLogger.SetOutput(infoSrc)

		//接口调用日志
		apiLogger = logrus.New()
		apiLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	Logger = &logger{
		OnFile: onFile,
		api:    apiLogger,
		error:  errLogger,
		info:   infoLogger,
		panic:  panicLogger,
	}
	return Logger, nil
}

func (logger *logger) SetApiFile(reqUri string) error {
	//构建目录和文件路径
	logFilePath, _ := os.Getwd()
	index := strings.IndexAny(reqUri, "?")
	if index != -1 {
		reqUri = reqUri[:index]
	}
	paths := strings.Split(reqUri, "/")
	fileName := logFilePath + "/log/api"
	dir := logFilePath + "/log/api"
	for index, path := range paths {
		fileName += "/" + path
		if index == len(paths)-1 {
			continue
		}
		dir += "/" + path
	}
	//生成日志目录
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		Logger.Error(err.Error())
		return err
	}

	//接口记录日志初始化
	apiSrc, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		Logger.error.Error(context.Background(), err.Error())
		return err
	}
	logger.api.SetOutput(apiSrc)

	return nil
}

func (logger *logger) Error(err ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if logger.OnFile {
		Logger.error.Errorln(err, file, line)
	}
	fmt.Println(err, file, line)
}

func (logger *logger) Api(err ...interface{}) {
	if logger.OnFile {
		Logger.api.Infoln(err)
	}
	fmt.Println(err)
}

func (logger *logger) Panic(err ...interface{}) {
	_, file, line, _ := runtime.Caller(4)
	if logger.OnFile {
		Logger.panic.Panicln(err, file, line)
	}
	fmt.Println(err, file, line)
}

func (logger *logger) Info(err ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if logger.OnFile {
		Logger.info.Infoln(err, file, line)
	}
	fmt.Println(err, file, line)
}
