package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func main()  {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 典型读取操作， 默认分区可以使用空字符串表示
	fmt.Println("App Mode", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path", cfg.Section("paths").Key("data").String())

	// 选值限制
	fmt.Println("Server Protocol:", cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 若读取的值不在候选列表内，则回退使用提供的默认值
	fmt.Println("Email Protocol:", cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 自动类型转换
	fmt.Printf("Port Number:(%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enfore Domain:(%[1]T) %[1]v\n", cfg.Section("server").Key("enfore_domain").MustBool(false))

	// 修改某个值后进行保存
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini.local")
}

