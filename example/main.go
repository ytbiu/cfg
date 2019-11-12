package main

import (
	"cfg"
	"log"
)

// 配置文件实体
type config struct {
	Dev *Dev `json:"dev1" yaml:"dev1"`
}

// 具体的 k-v结构
type Dev struct {
	Student *Student `json:"student" yaml:"student"`
}

// 配置项
type Student struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

var c config

func runExample() error {
	err := cfg.FromConsul(
		cfg.Addr("consul.exmarttech.com:80"),
		cfg.ConfigPath("dev1"),
		cfg.Result(&c),
		cfg.EncodeType(cfg.YamlEnc),
		cfg.Hooks(
			func() {
				log.Println("config name : ", c.Dev.Student.Name)
			}),
	)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := runExample();err != nil{
		log.Println(err)
		return
	}

	log.Println("first read config -- name: ", c.Dev.Student.Name)
	select {}
}
