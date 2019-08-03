package main

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
	"github.com/sparkstar/tuiwidget/loginwindow"
)

func main() {
	logo := `
 _   _       _ _       _    
| | | |     | (_)     | |   
| | | |_ __ | |_ _ __ | | __
| | | | '_ \| | | '_ \| |/ /
| |_| | |_) | | | | | |   < 
 \___/| .__/|_|_|_| |_|_|\_\
      | |                   
      |_|`

	loginWindow := loginwindow.New(logo)

	loginWindow.OnSubmit(func() {
		fmt.Println(loginWindow.GetText())
	})
	ui, _ := tui.New(loginWindow)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
