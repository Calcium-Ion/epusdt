package log

import (
	"fmt"
	"os"
	"time"

	"github.com/assimon/luuu/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger

func Init() {
	// Create file writer
	fileWriter := getLogWriter()
	// Create console writer
	consoleWriter := zapcore.AddSync(os.Stdout)

	// Combine both writers
	multiWriter := zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)

	encoder := getEncoder()
	core := zapcore.NewCore(encoder, multiWriter, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file := fmt.Sprintf("%s/log_%s.log",
		config.LogSavePath,
		time.Now().Format("20060102"))
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    config.LogMaxSize,
		MaxBackups: config.LogMaxBackups,
		MaxAge:     config.LogMaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
