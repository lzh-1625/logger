package main

import (
	"logger"
	"time"
)

func main() {
	logger := logger.NewDefaultLogger()
	logger.AddAdditionalInfo("a", "1")
	logger.Infow("routine1")
	go func() {
		logger.Infow("routine2")
	}()
	go func() {
		logger.AddAdditionalInfo("b", "2")
		logger.Infow("routine3")
	}()
	logger.DeleteAdditionalInfo(1)
	logger.Infow("routine1")
	time.Sleep(time.Second)
}

// INFO    test/test.go:11 routine1        {"a": "1"}
// INFO    test/test.go:20 routine1
// INFO    test/test.go:17 routine3        {"b": "2"}
// INFO    test/test.go:13 routine2
