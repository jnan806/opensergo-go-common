package main

import (
	"github.com/opensergo/opensergo-go-common/pkg/logging"
	"github.com/pkg/errors"
)

// when application has use logging-component in opensergo-go-common
// and import the other logging-component such as sentinel logger
// how to print logs with the sentinel logger through logging-component in opensergo-go-common
func main() {
	adaptor := NewLoggerAdaptor()
	logging.AppendLoggerSlice(adaptor)
	logging.Error(errors.New("errors.New"), "this is error log implement logging in sentinelLoggerAdaptor")
	logging.Info("this is info log implement logging in sentinelLoggerAdaptor")
	logging.Debug("this is debug log implement logging in sentinelLoggerAdaptor")
}
