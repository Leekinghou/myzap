package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

type LevelEnablerFunc func(Level) bool

type TeeOption struct {
	Out io.Writer
	LevelEnablerFunc
}

// NewTee 定义一个函数来构造 Logger 对象。
// @Param   tees []TeeOption 一个切片 每一个 TeeOption 对象对应一个日志级别
// @Param	opts ...Option	包含了一个 io.Writer 和一个 LevelEnablerFunc 函数 用来指定日志输出位置和日志级别，当满足定义的日志级别时将日志输出到指定位置
func NewTee(tees []TeeOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	for _, tee := range tees {
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.RFC3339TimeEncoder
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			zapcore.AddSync(tee.Out),
			zap.LevelEnablerFunc(tee.LevelEnablerFunc),
		)
		cores = append(cores, core)
	}
	return &Logger{l: zap.New(zapcore.NewTee(cores...), opts...)}
}
