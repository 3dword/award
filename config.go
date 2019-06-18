package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var Conf = struct {
	Award struct {
		StartTime string `toml:"startTime"`
		EndTime string `toml:"endTime"`
	}

	AwardMap map[string]int64

	Redis struct {
		Ip string `toml:"ip""`
		Port uint32 `toml:"port"`
		Network string `toml:"network"`
	}

	Mysql struct {
		Database string `toml:"database"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
}{}

/**
* 解析配置文件
 */
func parse() {
	if _, err := toml.DecodeFile("./conf/config.toml", &Conf); err != nil {
		fmt.Println("parse config error , ", err)
	}
 	fmt.Println("awardMap : ", Conf.AwardMap)
}
