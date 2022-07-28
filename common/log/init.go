package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
	"typora-qiniu-upload/common/config"
)

const FieldSeparator = " "

func InitLog() {
	logCfg := config.GetLogCfg()

	// Encoder
	encodeConfig := zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "name",
		CallerKey:        "caller",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      encodeLevel,
		EncodeTime:       encodeTime,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     encodeCaller,
		EncodeName:       nameEncoder,
		ConsoleSeparator: FieldSeparator,
	}
	encoder := zapcore.NewJSONEncoder(encodeConfig)

	// WriteSyncer
	file, _ := os.OpenFile(logCfg.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	zcore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(file)), zap.InfoLevel)

	listOpt := []zap.Option{
		zap.AddCaller(),
	}
	logger := zap.New(zcore, listOpt...)
	l := &Logger{zapLogger: logger}
	setLogger(l)
}

func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.CapitalString() + "]")
}

func encodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func nameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + loggerName + "]")
}
