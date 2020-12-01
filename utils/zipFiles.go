package utils

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func ZipFiles(filename string, files []string, oldForm, newForm string) error {

	// 创建新的文件
	newZipFile, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer newZipFile.Close()

	// 创建zip 写地址
	zipWriter := zip.NewWriter(newZipFile)

	defer zipWriter.Close()

	// 遍历文件列表
	for _, file := range files {
		// 打开文件
		zipFile, err := os.Open(file)

		if err != nil {
			return err
		}
		//
		defer zipFile.Close()
		// 获取文件状态
		info, err := zipFile.Stat()

		if err != nil {
			return err
		}
		// 文件头部信息
		header, err := zip.FileInfoHeader(info)

		if err != nil {
			return err
		}
		// 写名字
		header.Name = strings.Replace(file, oldForm, newForm, -1)
		//
		header.Method = zip.Deflate
		// 设置zip 写流的 header
		writer, err := zipWriter.CreateHeader(header)

		//
		if err != nil {
			return err
		}

		// 调用copy 方法。
		if _, err = io.Copy(writer, zipFile); err != nil {
			return err
		}
	}

	return nil
}
