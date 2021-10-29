package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//  logger
var logger *zap.Logger = nil

//配置变量默认值,读取配置文件进行修改
var appName = "logappdefault"                  // app标记
var logFilePath = "./logs/" + appName + ".log" // 保存路径
var logLevel = "debug"                         // 日志等级
var maxSizeDefault = 128                       // 每个日志文件保存的最大尺寸 单位：M
var maxBackupsDefault = 30                     // 日志文件最多保存多少个备份
var maxAgeDefault = 30                         // 文件最多保存多少天

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

//获取日志级别
func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

//NewLogger 初始化日志
func NewLogger(name string, logFile string, logLevelName string, maxSize int, maxBackups int, maxAge int) {

	if name == "" {
		name = appName
	}
	if logFile == "" {
		logFile = logFilePath
	}

	if logLevelName == "" {
		logLevelName = logLevel
	}

	if maxSize == 0 {
		maxSizeDefault = maxSize
	}
	if maxBackups == 0 {
		maxBackupsDefault = maxBackups
	}

	if maxAge == 0 {
		maxAgeDefault = maxAge
	}

	level := getLoggerLevel(logLevelName)
	hook := lumberjack.Logger{
		Filename:   logFile,           // 日志文件路径
		MaxSize:    maxSizeDefault,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackupsDefault, // 日志文件最多保存多少个备份
		MaxAge:     maxAgeDefault,     // 文件最多保存多少天
		Compress:   true,              // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     currentTimeEncoder,             // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("app", name))
	//因为对zap做了封装,在记录日志时,跳过本次封装,这样才能记录到真实的异常信息文件和行数,不然一直显示是本文件的方法行数
	skip := zap.AddCallerSkip(1)
	// 构造日志
	logger = zap.New(core, caller, development, filed, skip)

	//logger = logger.Sugar()
}

func currentTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	t = t.Local()
	enc.AppendString(t.Format("2006-01-02 15-04-05.000"))
}

//隔离引用
type LogField = zap.Field

func Debug(msg string, fields ...LogField) {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...LogField) {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...LogField) {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	logger.Warn(msg, fields...)
}

func Error(err error, fields ...LogField) {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	logger.Error(err.Error(), fields...)
}

//func DPanic(msg string, fields ...LogField) {
//	logger.DPanic(msg, fields...)
//}

func Panic(msg error, fields ...LogField) {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	logger.Panic(msg.Error(), fields...)
}

func Fatal(msg string, fields ...LogField) {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	logger.Fatal(msg, fields...)
}

func String(key string, val string) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.String(key, val)
}

func Int(key string, val int) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Int(key, val)
}
func Int32(key string, val int32) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Int32(key, val)
}
func Int64(key string, val int64) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Int64(key, val)
}
func Bool(key string, val bool) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Bool(key, val)
}

func Duration(key string, val time.Duration) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Duration(key, val)
}
func Time(key string, val time.Time) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Time(key, val)
}
func Float32(key string, val float32) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Float32(key, val)
}
func Float64(key string, val float64) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Float64(key, val)
}

func ByteString(key string, val []byte) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.ByteString(key, val)
}

func Uintptr(key string, val uintptr) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Uintptr(key, val)
}

func Any(key string, value interface{}) LogField {
	if logger == nil {
		NewLogger(appName, logFilePath, logLevel, maxSizeDefault, maxBackupsDefault, maxAgeDefault)
	}
	return zap.Any(key, value)
}
