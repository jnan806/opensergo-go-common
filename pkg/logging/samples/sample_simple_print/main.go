package main

import (
	"github.com/opensergo/opensergo-go-common/pkg/logging"
	"github.com/pkg/errors"
)

// use default logger to print log
func main() {
	printErrorStack()
	//printDefault()
}

func printErrorStack() {
	logging.NewConsoleLogger(logging.DebugLevel, logging.JsonFormat, true)
	logging.Error(errors.New("errors.New"), "this is error log in printErrorStack()")
	logging.Info("this is info log in printErrorStack()")
	logging.Debug("this is debug log in printErrorStack()")
}

func printDefault() {
	logging.NewDefaultConsoleLogger(logging.DebugLevel)
	logging.Error(errors.New("errors.New"), "this is error log in printDefault()")
	logging.Info("this is info log in printDefault()")
	logging.Debug("this is debug log in printDefault()")
}
