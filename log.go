package main

import (
	"fmt"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

func LogErrorf(format string, a ...interface{}) {
	cqp.AddLog(cqp.Error, "MC版本检查", fmt.Sprintf(format, a...))
}

func LogInfof(format string, a ...interface{}) {
	cqp.AddLog(cqp.Info, "MC版本检查", fmt.Sprintf(format, a...))
}
