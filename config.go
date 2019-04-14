package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var conf struct {
	Award struct {
		A int64 `toml:"A"`
		B int64 `toml:"B"`
		C int64 `toml:"C"`
		StartTime string `toml:"StartTime"`
		EndTime string `toml:"EndTime"`
	}
}

/**
* 解析配置文件
 */
func parse() {
	if _, err := toml.DecodeFile("./conf/config.toml", &conf); err != nil {
		fmt.Println("parse config error , ", err)
	}
}
