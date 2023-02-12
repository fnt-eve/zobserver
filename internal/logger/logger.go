package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//nolint:revive
type LoggerConfig struct {
	LogLevel  string `default:"debug" split_words:"true"`
	LogFormat string `default:"devel" split_words:"true"`
}

func (c *LoggerConfig) InitLogger() (*zap.Logger, error) {
	logLevel := strings.ToLower(c.LogLevel)

	var aLevel zap.AtomicLevel
	aLevel, err := zap.ParseAtomicLevel(logLevel)
	// Default log level to debug if parsing failed
	if err != nil {
		aLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	var cfg zap.Config
	// Set formatting to prod if prod, devel otherwise
	if c.LogFormat == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	// Always set timestamp formatting to ISO8601
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.Level = aLevel

	return cfg.Build()
}
