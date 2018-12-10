package manager

import (
	"github.com/isayme/go-email/app/conf"
	"github.com/isayme/go-email/app/email"
)

var globalManager Manager

// Manager global context manager
type Manager struct {
	Sender email.Sender
}

// Init init
func Init(config *conf.Config) {
	if config.Sender.MailGun.Domain != "" {
		globalManager.Sender = email.NewMailGun(&config.Sender.MailGun)
	}
}

// Get return global manager
func Get() *Manager {
	return &globalManager
}
