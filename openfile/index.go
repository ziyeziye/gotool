package openfile

import (
	"os"
)

type FileUtils struct {
}

// Exists 判断所给路径文件或文件夹是否存在
func (FileUtils) Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func (FileUtils) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func (f FileUtils) IsFile(path string) bool {
	return !f.IsDir(path)
}

// RemoveFile 删除文件，参数文件路径
func (f FileUtils) RemoveFile(path string) error {
	//删除文件
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
