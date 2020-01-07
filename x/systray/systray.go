package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

// systray 是一款golang 类库，可以提供不同平台下面的系统栏的菜单功能
func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac.
	mQuit.SetIcon(icon.Data)
}

func onExit() {
	// clean up here
}
