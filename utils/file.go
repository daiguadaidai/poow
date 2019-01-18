package utils

import (
	"fmt"
	"github.com/cihub/seelog"
	"io"
	"os"
	"path/filepath"
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

// 拷贝文件
func FileCopy(src, std string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	stdFile, err := os.Create(std)
	if err != nil {
		return err
	}

	if _, err = io.Copy(srcFile, stdFile); err != nil {
		return err
	}

	return nil
}

type TailData struct {
	Info string `json:"info"`
	End  int64  `json:"end"`
}

const DEFAULT_TAIL_SIZE = 524288

// 读取文件数据
// return data(数据), end(本次获取结束位点), error(错误)
func TailFile(filePath string, start int64, size int64) (*TailData, error) {
	tailData := new(TailData)
	var offset int64
	if size == 0 {
		size = DEFAULT_TAIL_SIZE
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return tailData, err
	}
	f, err := os.Open(filePath)
	if err != nil {
		return tailData, err
	}
	defer f.Close()

	var b []byte
	if start == 0 { // 没有指定文件开始字节位点
		if fileInfo.Size()-size > 0 {
			start = fileInfo.Size() - size
			offset = fileInfo.Size() - start
		} else {
			size = fileInfo.Size()
			offset = 0
		}
	} else { // 有指定开始位点
		if fileInfo.Size() <= start { // 没有数据
			tailData.End = fileInfo.Size()
			return tailData, nil
		}

		if fileInfo.Size()-start <= size { // 获取指定大小数据, 能到文件最后
			offset = start
			size = fileInfo.Size() - start
		} else { // 不能到文件最后
			offset = start
		}
	}
	b = make([]byte, size)
	n, err := f.ReadAt(b, offset)

	tailData.Info = string(b)
	tailData.End = start + int64(n)
	return tailData, nil
}
