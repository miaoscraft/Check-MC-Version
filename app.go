package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"log"
	"net/http"
	"time"
)

//go:generate cqcfg -c .
// cqp: 名称: MC版本更新推送
// cqp: 版本: 1.0.0:0
// cqp: 作者: Tnze
// cqp: 简介: 自动检查Minecraft版本更新
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "cn.miaoscraft.mc-checker"

	ctx, cancel := context.WithCancel(context.Background())
	cqp.Enable = func() int32 {
		go checkLoop(ctx)
		return 0
	}

	cqp.Disable = func() int32 {
		cancel()
		return 0
	}
}

func checkLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			list, err := check()
			if err != nil {
				cqp.AddLog(cqp.Error, "MC更新通知", fmt.Sprintf("检查更新出错: %v", err))
			} else {
				notice(list)
			}
		}
	}
}

func notice(list versions) {
	for _, v := range list.Versions {
		if v.ID == list.Latest.Release {
			log.Println("最新版本：", v.ID, "发布日期：", v.Time.Format("2006-01-02"))
		}
		if v.ID == list.Latest.Snapshot {
			log.Println("最新快照：", v.ID, "发布日期：", v.Time.Format("2006-01-02"))
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
	Versions []struct {
		ID          string
		Type        string
		URL         string
		Time        time.Time
		ReleaseTime time.Time
	}
}
