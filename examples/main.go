package main

import (
	log "myzip"
)

func main() {
	// 之前
	//config := zap.NewProductionConfig()
	//config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	//logger, _ := config.Build()
	//defer logger.Sync()
	//logger.Error("log error")

	// 现在
	//defer log.Sync()
	//log.Info("Info msg")
	//
	//log.SetLevel(log.ErrorLevel)
	//log.Info("Info msg")
	//log.Error("Error msg")

	//logger := log.New(os.Stderr, log.ErrorLevel)
	//log.ReplaceDefault(logger)
	//defer log.Sync()
	//log.Info("Info msg")
	//log.Error("Error msg")

	// 日志输出选项
	// log.AddCaller() 选项用来开启记录，log.AddCallerSkip(1) 用来设置通过调用栈获取文件名和行号时跳过的调用深度
	//logger := log.New(os.Stderr, log.InfoLevel, log.AddCaller(), log.AddCallerSkip(1))
	//defer logger.Sync()
	//logger.Info("Info msg")

	// 支持将不同级别日志输出到不同位置
	// 在 zap 中可以通过 zapcore.NewTee() 实现，它返回一个切片 []zapcore.Core，这样每一个 zapcore.Core 对应一种日志级别，就能够实现将不同级别日志输出到不同位置了。
	//file, _ := os.OpenFile("test-warn.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//tees := []log.TeeOption{
	//	{
	//		Out: os.Stdout,
	//		LevelEnablerFunc: func(level log.Level) bool {
	//			return level == log.InfoLevel
	//		},
	//	},
	//	{
	//		Out: file,
	//		LevelEnablerFunc: func(level log.Level) bool {
	//			return level == log.WarnLevel
	//		},
	//	},
	//}
	//logger := log.NewTee(tees)
	//defer logger.Sync()
	//
	//logger.Info("Info tee msg")
	//logger.Warn("Warn tee msg")
	//logger.Error("Error tee msg") // 不会输出

	// 日志轮转
	tees := []log.TeeOption{
		{
			Out: log.NewProductionRotateBySize("rotate-by-size.log"),
			LevelEnablerFunc: log.LevelEnablerFunc(func(level log.Level) bool {
				return level < log.WarnLevel
			}),
		},
		{
			Out: log.NewProductionRotateByTime("rotate-by-time.log"),
			LevelEnablerFunc: log.LevelEnablerFunc(func(level log.Level) bool {
				return level >= log.WarnLevel
			}),
		},
	}
	lts := log.NewTee(tees)
	defer lts.Sync()

	lts.Debug("Debug msg")
	lts.Info("Info msg")
	lts.Warn("Warn msg")
	lts.Error("Error msg")
}
