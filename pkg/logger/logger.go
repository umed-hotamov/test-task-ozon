package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}

	return true
}

func createDir(path string) error {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	return nil
}

func createLogDir(path string) error {
	ok := pathExists(path)
	if ok {
		return nil
	}

	return createDir(path)
}

func NewLogger(level string) (*zap.Logger, error) {
  setLevel, exists := levelMapping[level]
  if !exists {
    log.Fatalf("Invalid log level")
  }

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writer, setLevel),
	)

	return zap.New(core, zap.AddCaller()), nil
}
