package main

import (
	"io/ioutil"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	config.Init()
	database.Init()
	servers.Init()

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon(config.HttpDirectory + "/favicon.ico"))
	systray.SetTooltip("Go Avito Parser")
	trayCP := systray.AddMenuItem("Open control panel", "")
	systray.AddSeparator()
	trayQuit := systray.AddMenuItem("Quit", "")
	go func() {
		for {
			select {
			case <-trayCP.ClickedCh:
				_ = open.Run("http://" + config.BindAddress)
			case <-trayQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return b
}
