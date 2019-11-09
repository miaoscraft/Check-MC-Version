package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"strings"
)

func sendMsg(v version) {
	var sb strings.Builder
	err := temp.Execute(&sb, v)
	if err != nil {
		LogErrorf("执行模版出错: %v", err)
	}

	msg := sb.String()
	LogInfof("%s", msg)

	for _, gid := range conf.GroupID {
		cqp.SendGroupMsg(int64(gid), msg)
	}

}
