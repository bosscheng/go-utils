package utils

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	//
	_, err := os.Stat(filename)
	//
	return err == nil || os.IsExist(err)
}

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// read file line
// 读取文件内容，返回字符串数组。
func ReadFileLines(filename string) ([]string, error) {
	var lines []string
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return lines, err
	}

	defer file.Close()

	bio := bufio.NewReader(file)

	for {
		var line []byte
		line, _, err := bio.ReadLine()

		if err != nil {
			if err == io.EOF {
				file.Close()
				return lines, nil
			}
			return nil, nil
		}
		lines = append(lines, string(line))
	}

	return lines, nil
}

// current dir
// 返回
func CurrentDir(path ...string) string {
	_, currentFilePath, _, _ := runtime.Caller(1)

	if len(path) == 0 {
		return filepath.Dir(currentFilePath)
	}

	return filepath.Join(filepath.Dir(currentFilePath), filepath.Join(path...))
}

// create dir 批量创建文件夹
func CreateDir(dirs ...string) error {
	for _, v := range dirs {
		exist, err := PathExist(v)

		if err != nil {
			return err
		}

		if !exist {
			err = os.MkdirAll(v, os.ModePerm)

			if err != nil {
				return err
			}

		}
	}
	return nil
}
