package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/gwyy/img-cloud-update/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	now := time.Now()
	infoLogFileName := fmt.Sprintf("%s/info/%04d-%02d-%02d.log", global.Conf.Log.Path, now.Year(), now.Month(), now.Day())
	errorLogFileName := fmt.Sprintf("%s/error/%04d-%02d-%02d.log", global.Conf.Log.Path, now.Year(), now.Month(), now.Day())
	var coreArr []zapcore.Core

	// 获取编码器
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 指定时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // ，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	////encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        // 显示完整文件路径

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "file",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		//EncodeTime: zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		//EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		//	enc.AppendInt64(int64(d) / 1000000)
		//},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		//EncodeCaller: zapcore.FullCallerEncoder,
		//EncodeName:       nil,
		//ConsoleSeparator: "",
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zap.ErrorLevel && level >= zap.DebugLevel
	})

	// 当yml配置中的等级大于Error时，lowPriority级别日志停止记录
	if global.Conf.Log.Level >= 2 {
		lowPriority = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return false
		})
	}

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoLogFileName,            //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    global.Conf.Log.MaxSize,    //文件大小限制,单位MB
		MaxAge:     global.Conf.Log.MaxAge,     //日志文件保留天数
		MaxBackups: global.Conf.Log.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   global.Conf.Log.Compress, //是否压缩处理
	})
	// 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorLogFileName,           //日志文件存放目录
		MaxSize:    global.Conf.Log.MaxSize,    //文件大小限制,单位MB
		MaxAge:     global.Conf.Log.MaxAge,     //日志文件保留天数
		MaxBackups: global.Conf.Log.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   global.Conf.Log.Compress, //是否压缩处理
	})
	// 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	global.Log = logger.Sugar()
	global.Log.Info("初始化zap日志完成!")

}
