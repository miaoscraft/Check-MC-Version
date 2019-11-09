// [试验性]配置文件管理器
// 指定文件名获取文件，如果文件不存在则用默认值创建文件后返回
package resfile

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

// GetFile 返回一个只读文件，当文件不存在时用默认数据创建
// 当filename是相对路径时，认为该路径基于插件的AppDir
func GetFile(filename string, defaultdata []byte) (*os.File, error) {
	// 计算绝对路径
	if !filepath.IsAbs(filename) {
		filename = filepath.Join(cqp.GetAppDir(), filename)
	}
	// 打开文件
	f, err := os.Open(filename)
	// 文件不存在时自动创建
	if os.IsNotExist(err) {
		if err := ioutil.WriteFile(filename, defaultdata, 0666); err != nil {
			return nil, err
		}

		f, err = os.Open(filename)
	}
	if err != nil {
		return nil, err
	}

	return f, nil
}
