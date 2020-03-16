package main

import (
	"context"
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//go:generate cqcfg -c .
// cqp: 名称: MC版本更新推送
// cqp: 版本: 1.0.0:0
// cqp: 作者: Tnze
// cqp: 简介: 自动检查Minecraft版本更新
func main() {}

var conf config             // 插件设置
var temp *template.Template // 通知模版
var latestRelease, latestSnapshot string

func init() {
	cqp.AppID = "cn.miaoscraft.mc-checker"

	ctx, cancel := context.WithCancel(context.Background())
	cqp.Enable = func() int32 {
		var err error
		conf, err = getConfig()
		if err != nil {
			LogErrorf("读取配置出错: %v", err)
			return 1
		}

		temp, err = template.New("update").Parse(conf.Template)
		if err != nil {
			LogErrorf("解析模版出错: %v", err)
			return 1
		}

		go checkLoop(ctx)
		return 0
	}

	cqp.Disable = func() int32 {
		cancel()
		return 0
	}
}

func checkLoop(ctx context.Context) {
	//立刻检查并记录当前MC版本号
	list, err := check()
	if err != nil {
		LogErrorf("检查更新出错: %v", err)
	} else {
		latestRelease = list.Latest.Release
		latestSnapshot = list.Latest.Snapshot
	}

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			list, err := check()
			if err != nil {
				LogErrorf("检查更新出错: %v", err)
			} else {
				notice(list)
			}
		}
	}
}

func notice(list versions) {
	for _, v := range list.Versions {
		if v.ID == list.Latest.Release && v.ID != latestRelease {
			sendMsg(v)
			latestRelease = list.Latest.Release
		}
		if v.ID == list.Latest.Snapshot && v.ID != latestSnapshot {
			sendMsg(v)
			latestSnapshot = list.Latest.Snapshot
		}
	}
}

func check() (list versions, err error) {
	resp, err := http.Get("https://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return list, err
	}

	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return list, err
	}

	return list, nil
}

type versions struct {
	Latest struct {
		Release, Snapshot string
	}
	Versions []version
}

type version struct {
	ID          string
	Type        string
	URL         string
	Time        time.Time
	ReleaseTime time.Time
}
