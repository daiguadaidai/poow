package utils

import (
	"fmt"
	"github.com/cihub/seelog"
	"os"
	"path/filepath"
	"time"
)

// 文件/目录 是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 获取绝对路径
func AbsPath(_path string) string {
	absPath, err := filepath.Abs(_path)
	if err != nil {
		return _path
	}

	return absPath
}

// 创建一个目录
func CreateDir(_path string) error {
	err := os.MkdirAll(_path, os.ModePerm)

	return err
}

/* 检测和创建路径
Params:
    _path: 路径
    _pathUsage: 是做什么用的
*/
func CheckAndCreatePath(_path string, _pathUsage string) error {
	if exists, err := PathExists(_path); err != nil {
		return fmt.Errorf("%v, 检测失败. %v", _pathUsage, err)
	} else {
		if !exists {
			seelog.Warnf("%v, 不存在: %s(%s)",
				_pathUsage, _path, AbsPath(_path))

			err1 := CreateDir(_path)
			if err1 != nil {
				return fmt.Errorf("%v失败: %s(%s). %v",
					_pathUsage, _path, AbsPath(_path), err1)
			}
			seelog.Warnf("创建%v成功: %s(%s)",
				_pathUsage, _path, AbsPath(_path))
		}
	}

	return nil
}

// 生成一个全局唯一的文件名
func CreateUUIDFileName() string {
	t := time.Now()
	fileName := fmt.Sprintf("%v", t.Format("20060102150405123456"))

	return fileName
}

func ChmodFile(_path string) error {
	return os.Chmod(_path, os.ModePerm)
}

// 判断文件是否可执行
func FileIsExecutable(_path string) (bool, error) {
	fileInfo, err := os.Stat(_path)
	if err != nil {
		return false, err
	}

	mode := fileInfo.Mode()
	perm := mode.Perm()
	flag := perm & os.FileMode(73)

	if uint32(flag) == uint32(73) {
		return true, nil
	}

	return false, nil
}

// 删除文件
func RemoveFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

// 获取文件大小
func FileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
