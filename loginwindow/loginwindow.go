package loginwindow

import (
	"github.com/marcusolsson/tui-go"
	entry "github.com/sparkstar/tuiwidget/labeledentry"
)

type LoginWindow struct {
	*tui.Box
	logo string

	username *entry.LabeledEntry
	password *entry.LabeledEntry

	onSubmit func()
}

func createLabeledEntry(name string, mode tui.EchoMode) *entry.LabeledEntry {
	l := entry.New(name, mode)
	l.SetBorder(true)
	return l
}

func makeCentered(widget tui.Widget) *tui.Box {
	return tui.NewHBox(tui.NewSpacer(), widget, tui.NewSpacer())
}

func New(logo string) *LoginWindow {
	logolabel := tui.NewLabel(logo)
	logobox := tui.NewHBox(tui.NewSpacer(), logolabel, tui.NewSpacer())

	user := createLabeledEntry("Username", tui.EchoModeNormal)
	pass := createLabeledEntry("Password", tui.EchoModePassword)

	button := tui.NewButton("Login")
	centeredbtn := makeCentered(button)

	tui.DefaultFocusChain.Set(user, pass, button)

	window := tui.NewVBox(logobox, tui.NewSpacer(), user, pass, centeredbtn)
	window.SetBorder(true)

	return &LoginWindow{
		Box:  window,
		logo: logo,

		username: user,
		password: pass,
	}
}

func (lw *LoginWindow) OnKeyEvent(ev tui.KeyEvent) {
	if ev.Key == tui.KeyEnter {
		lw.onSubmit()
	}

	lw.Box.OnKeyEvent(ev)
}

func (lw LoginWindow) GetText() (string, string) {
	return lw.username.Text(), lw.password.Text()
}

func (lw *LoginWindow) OnSubmit(fn func()) {
	lw.onSubmit = fn
}
