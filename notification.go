package notification

import (
	"bytes"
	"os/exec"
	"text/template"
)

var OSXTemplate = template.Must(template.New("osx").Parse(`display notification "{{.Message}}" with title "{{.Title}}" subtitle "{{.SubTitle}}" sound name "{{.Sound}}"`))

type Notification struct {
	Message  string
	Title    string
	SubTitle string
	Sound    string
}

func Message(msg string) *Notification {
	return &Notification{Message: msg}
}

func (n *Notification) WithTitle(title string) *Notification {
	n.Title = title
	return n
}

func (n *Notification) WithSubTitle(subTitle string) *Notification {
	n.SubTitle = subTitle
	return n
}

func (n *Notification) WithSound(sound string) *Notification {
	n.Sound = sound
	return n
}

func (n *Notification) Push() {
	var buf bytes.Buffer
	OSXTemplate.Execute(&buf, n)

	exec.Command("osascript", "-e", buf.String()).Run()
}
