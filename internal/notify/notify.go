package notify

import (
	"github.com/martinlindhe/notify"
)

type Notifier interface {
	Notify(appName, title, text, iconPath string)
}

type SystemNotifier struct{}

func (n SystemNotifier) Notify(appName, title, text, iconPath string) {
	notify.Notify(appName, title, text, iconPath)
}

func Notify(n Notifier, title, message string) error {
	n.Notify("Bug Out Monitor", "notice", title+" -- "+message, "path/to/icon.png")
	return nil
}
