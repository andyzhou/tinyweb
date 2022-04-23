package main

import (
	"github.com/andyzhou/tinyweb"
)

//inter macro define
const (
	serverPort = 8080
	tplDir = "./tpl"
)

func main() {
	//get web instance
	web := tinyweb.GetWeb()

	//init web app
	app := web.GetWebApp()

	//base setup
	app.SetTplPath(tplDir)

	//register sub app
	app.RegisterSubApp("/", NewSubApp())

	//start app
	app.Start(serverPort)
}

//