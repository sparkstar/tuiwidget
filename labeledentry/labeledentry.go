package labeledentry

import (
	"github.com/marcusolsson/tui-go"
)

const (
	DefaultSizeTextLength = 20
)

type LabeledEntry struct {
	*tui.Box
	tag         *tui.Label
	input       *tui.Entry
	maxSizeText int
}

func New(label string, inputmode tui.EchoMode) *LabeledEntry {
	l := tui.NewLabel(label)
	p := tui.NewPadder(1, 0, l)
	e := tui.NewEntry()
	b := tui.NewHBox(p, e)

	l.SetSizePolicy(tui.Maximum, tui.Preferred)
	e.SetSizePolicy(tui.Expanding, tui.Preferred)
	e.SetEchoMode(inputmode)

	r := &LabeledEntry{
		Box:         b,
		tag:         l,
		input:       e,
		maxSizeText: DefaultSizeTextLength,
	}

	e.OnChanged(func(entry *tui.Entry) {
		if len(entry.Text()) > r.maxSizeText {
			e.SetText(e.Text()[:r.maxSizeText])
		}
	})

	return r
}

func (l *LabeledEntry) SetMaxTextLength(length int) {
	l.maxSizeText = length
}

func (l *LabeledEntry) SetFocused(focused bool) {
	l.input.SetFocused((focused))
}

func (l LabeledEntry) Text() string {
	return l.input.Text()
}
