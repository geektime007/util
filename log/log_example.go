// +build ignore

package main

import (
	"github.com/geektime007/util/log"
)

var logger = log.Get("ExampleName")

func main() {
	logger.SetLevel(log.INFO)
	logger.SetNameLength(20)
	logger.SetMeta("geek_id", 531)
	logger.SetMeta("geek-name", "geek-048")
	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Warnf("This is a number %v", 1)
	logger.Warnf("This is a number %v", 1)
	logger.Warnf("This is a number %v", 1)
}
