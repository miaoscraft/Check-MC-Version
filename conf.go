package main

import (
	"github.com/BurntSushi/toml"
	"github.com/miaoscraft/Check-MC-Version/resfile"
	"time"
)

type config struct {
	Frequency duration
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
Frequency = 10m
`
