package notify

import (
	"testing"
)

type MockNotifier struct {
	appName, title, text, iconPath string
}

func (n *MockNotifier) Notify(appName, title, text, iconPath string) {
	n.appName, n.title, n.text, n.iconPath = appName, title, text, iconPath
}

func TestNotify(t *testing.T) {
	mock := &MockNotifier{}
	title := "Test Title"
	message := "Test Message"
	err := Notify(mock, title, message)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mock.title != "notice" || mock.text != title+" -- "+message {
		t.Errorf("Notify() = %q, %q; want %q, %q", mock.title, mock.text, "notice", title+" -- "+message)
	}
}
