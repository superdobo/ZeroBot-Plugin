package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"

	_ "github.com/Yiwen-Chan/ZeroBot-Plugin/chat"
	_ "github.com/Yiwen-Chan/ZeroBot-Plugin/github"
	_ "github.com/Yiwen-Chan/ZeroBot-Plugin/manager"
	_ "github.com/Yiwen-Chan/ZeroBot-Plugin/music"
	_ "github.com/Yiwen-Chan/ZeroBot-Plugin/runcode"
	//_ "github.com/Yiwen-Chan/ZeroBot-Plugin/setutime"
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	fmt.Printf(`
====================[ZeroBot-Plugin]====================
* OneBot + ZeroBot + Golang
* Copyright © 2020-2021 Kanri, All Rights Reserved
* Project: https://github.com/Yiwen-Chan/ZeroBot-Plugin
========================================================
`)
	zero.Run(zero.Config{
		NickName:      []string{"椛椛"},
		CommandPrefix: "/",
		SuperUsers:    []string{"825111790", "213864964"}, // 必须修改，否则无权限
		Driver: []zero.Driver{
			driver.NewWebSocketClient("127.0.0.1", "6700", ""),
		},
	})
	select {}
}