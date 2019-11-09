package main

import (
	"github.com/BurntSushi/toml"
	"github.com/miaoscraft/Check-MC-Version/resfile"
	"time"
)

type config struct {
	Frequency duration
	GroupID   []uint64
	Template  string
}

type duration struct{ time.Duration }

func (d *duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return
}

func getConfig() (config, error) {
	var c config
	f, err := resfile.GetFile("config.toml", []byte(defaultConfig))
	if err != nil {
		return c, err
	}

	_, err = toml.DecodeReader(f, &c)
	return c, err
}

const defaultConfig = `# 配置文件
Frequency = "10m" 	# 检查频率
GroupID = [123456] 	# 群号(支持多个群)
Template = '''
Mojang于{{ .time.Format "2006-01-02 15:04:05" }}发布了
Minecraft {{ .ID }}{{ if eq .Type "snapshot" }}快照{{ end }}
'''	# 提醒模版
`
