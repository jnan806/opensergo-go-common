package main

import (
	"github.com/opensergo/opensergo-go-common/pkg/logging"
	"github.com/pkg/errors"
)

// when application has use logging-component Self-contained
// and import logging-component in opensergo-go-common at the same time
// how to print logs with the format from logging-component in opensergo-go-common through logging-component Self-contained
func main() {
	logger := logging.NewConsoleLogger(logging.DebugLevel, logging.JsonFormat, true)
	sentinelLogger := NewSentinelLogger(logger, logging.JsonFormat, true)
	sentinelLogger.Error(errors.New("errors.New"), "this is error log use logging in sentinelLogger")
	sentinelLogger.Info("this is info log use logging in sentinelLogger")
	sentinelLogger.Debug("this is debug log use logging in sentinelLogger")
}
