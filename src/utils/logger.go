package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface{}

type logger struct {
	Logger *zap.Logger
	Suger  *zap.SugaredLogger
}

var (
	debugEncoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	prodEncoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
)

func NewLogger() Logger {
	var core = zapcore.NewTee(
		zapcore.NewCore(debugEncoder, os.Stdout, zap.DebugLevel),
	)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &logger{
		Logger: zapLogger,
		Suger:  zapLogger.Sugar(),
	}
}

func (l *logger) Debug(msg string) {
	defer l.Logger.Sync()
	l.Logger.Debug(msg)
}
func (l *logger) Debugf(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Debugf(template, args...)
}
func (l *logger) Debugw(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Debugw(template, keysAndValues...)
}

func (l *logger) Info(msg string) {
	defer l.Logger.Sync()
	l.Logger.Debug(msg)
}
func (l *logger) Infof(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Debugf(template, args...)
}
func (l *logger) Infow(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Debugw(template, keysAndValues...)
}
