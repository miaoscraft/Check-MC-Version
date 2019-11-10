# \[QQ机器人\]MC版本更新推送
定时检查Minecraft版本更新，当有更新时在群内发送通知

## 酷Q论坛(cqp.cc)发布贴：
https://cqp.cc/t/45770
(内附下载链接)
## 配置(必须)
请在酷Q目录\data\app\cn.miaoscraft.mc-checker\config.toml查找并编辑配置文件。
```
## MC版本更新推送姬配置文件

## 检查频率
Frequency = "10m"

## 群号(支持多个群)
GroupID = [123456]

## 提醒模版
Template = '''
{{ .Time.Format "2006-01-02 15:04:05" }}
Minecraft {{ .ID }}{{ if eq .Type "snapshot" }}快照版{{ else if eq .Type "release" }}正式版{{ end }}发布了！！'''
```
普通用户只需修改GroupID那行的123456为你想发送推送的群号，如果有多个群可写为[123456, 234567, 345678]这样。

## 高级配置
**Frequency**检查更新的时间间隔，toml配置文件中为字符串，
支持的写法请参考https://golang.google.cn/pkg/time/#ParseDuration

> 其实这个很灵活，例如每10分钟就写"10m"，30秒就是"30s"，五分半就是"5m30s"。

**Template**推送消息模版，如要修改请参考https://golang.google.cn/pkg/text/template/
支持的字段有

字段        |类型     |备注
------------|---------|------
.ID         |string   |版本名，例如1.14.4或19w45b
.Type       |string   |发布类型，可能为snapshot或release
.URL        |string   |版本清单数据链接（请不要在意这个）
.Time       |time.Time|版本时间
.ReleaseTime|time.Time|版本发布时间

> 这个模版语法要弄懂可要花点功夫哟，但是它的功能确实很强大！

## Q&A
> 这个推送的是什么版本的MC呀？  
是Minecraft: Java Edition。
