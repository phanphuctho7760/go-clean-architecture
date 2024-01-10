package logs

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func newZapLogger(logLevel string) (logger LogItf, shutdownFunc func()) {
	logLevelMap := map[string]zapcore.Level{
		"debug":  zap.DebugLevel,
		"info":   zap.InfoLevel,
		"warn":   zap.WarnLevel,
		"error":  zap.ErrorLevel,
		"dpanic": zap.DPanicLevel,
		"panic":  zap.PanicLevel,
		"fatal":  zap.FatalLevel,
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.Lock(zapcore.AddSync(os.Stdout)),
		logLevelMap[logLevel],
	)

	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return ZapLogger{
			logger: l,
		}, func() {
			//			This is unnecessary when not write to file
			//			err := logger.Sync()
			//			if err != nil {
			//				log.Printf("%sError to sync log zap error: %s at %s%s\n", constants.ANSIColorRed, err.Error(), helpers.GetCallerLocationSkip(1), constants.ANSIColorWhite)
			//				return
			//			}
		}
}

func (receiver ZapLogger) Debug(args ...interface{}) {
	receiver.logger.Debug(makeMessageLog(args), zap.Any("args", args))
}

func (receiver ZapLogger) Info(args ...interface{}) {
	receiver.logger.Info(makeMessageLog(args), zap.Any("args", args))
}

func (receiver ZapLogger) Warn(args ...interface{}) {
	receiver.logger.Warn(makeMessageLog(args), zap.Any("args", args))
}

func (receiver ZapLogger) Error(args ...interface{}) {
	receiver.logger.Error(makeMessageLog(args), zap.Any("args", args))
}

func (receiver ZapLogger) Fatal(args ...interface{}) {
	receiver.logger.Fatal(makeMessageLog(args), zap.Any("args", args))
}

func (receiver ZapLogger) Panic(args ...interface{}) {
	receiver.logger.Panic(makeMessageLog(args), zap.Any("args", args))
}

func (receiver ZapLogger) Debugf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	receiver.logger.Debug(message, zap.Any("args", args))
}

func (receiver ZapLogger) Infof(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	receiver.logger.Info(message, zap.Any("args", args))
}

func (receiver ZapLogger) Warnf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	receiver.logger.Warn(message, zap.Any("args", args))
}

func (receiver ZapLogger) Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	receiver.logger.Error(message, zap.Any("args", args))
}

func (receiver ZapLogger) Fatalf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	receiver.logger.Fatal(message, zap.Any("args", args))
}

func (receiver ZapLogger) Panicf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	receiver.logger.Panic(message, zap.Any("args", args))
}

func makeMessageLog(args []interface{}) string {
	if len(args) > 0 {
		return fmt.Sprint(args[0])
	}
	return "No message available"
}
