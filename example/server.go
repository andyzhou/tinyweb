package main

import (
	"github.com/andyzhou/tinyweb/face"
	"github.com/andyzhou/tinyweb/iface"
)

//inter macro define
const (
	serverPort = 8080
	tplDir = "./tpl"
)

func main() {
	var (
		app iface.IWebApp
	)

	//init web app
	app = face.NewWebApp()

	//base setup
	app.SetTplPath(tplDir)

	//register sub app
	app.RegisterSubApp("/", NewSubApp())

	//start app
	app.Start(serverPort)
}

//