package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath) // 返回文件信息结构描述文件
	switch {
	case os.IsNotExist(err): // 文件/目录是否存在
		mkDir()
	case os.IsPermission(err): // 权限是否满足
		log.Fatalf("Permission: %v", err)
	}
	// 调用文件，支持传入文件名称、指定模式调用文件、文件权限，返回的文件方法可以用于I/O
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile: %v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd() // 返回与当前目录对应的根路径名
	// os.ModePerm 定义文件权限为 0777
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm) // 创建对应的目录以及所需的子目录，成功返回 nil，失败返回 error
	if err != nil {
		panic(err)
	}
}
