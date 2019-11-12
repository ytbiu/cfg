package main

import (
	"log"
	conf "github.com/ytbiu/cfg"
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

	err := conf.FromConsul(
		conf.Addr("consul.exmarttech.com:80"),
		conf.ConfigPath("dev1"),
		conf.Result(&c),
		conf.EncodeType(conf.YamlEnc),
		conf.Hooks(
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
