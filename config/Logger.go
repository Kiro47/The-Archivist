package config

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// ZLog Global logger instance
var ZLog *zap.Logger

// InitLogger Initialize Zap logger with Lumberjack log file
func InitLogger() {
	/*
	I could break this up into multiple functions but it really doesn't help
	at all with readability and none of it is used twice.
	 */
	// Lumberjack log rotation config
	loggerJack := lumberjack.Logger{
		Filename:   Log.Path,
		MaxSize:    Log.MaxSize,
		MaxAge:     Log.MaxAge,
		MaxBackups: Log.MaxBackups,
		LocalTime:  false,
		Compress:   Log.Compress,
	}
	// Zap config init
	atomicLevel := zap.NewAtomicLevel()
	if Log.Level == "debug" {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "file",
		FunctionKey:      "func",
		StacktraceKey:    "trace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "",
	}
	var writers = []zapcore.WriteSyncer {
		zapcore.AddSync(&loggerJack),
		zapcore.AddSync(os.Stdout),
	}
	zcore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		atomicLevel,
		)
	// Open dev options in debug
	if  Log.Level == "debug" {
		caller := zap.AddCaller()
		dev := zap.Development()
		fields := zap.Fields(zap.String("appName", "Archivist"))
		ZLog = zap.New(zcore, caller, dev, fields)
	} else {
		ZLog = zap.New(zcore)
	}
	ZLog.Info("Successfully initialized zap logger")
}
