package email

import (
	"github.com/isayme/go-email/app/conf"
	mailgun "github.com/mailgun/mailgun-go"
)

type MailGun struct {
	config *conf.MailGun
	client *mailgun.MailgunImpl
}

func NewMailGun(config *conf.MailGun) Sender {
	client := mailgun.NewMailgun(config.Domain, config.APIKey)
	return &MailGun{
		config: config,
		client: client,
	}
}

func (mg *MailGun) Send(message *Message) (*MessageID, error) {
	m := mailgun.NewMessage(message.From, message.Subject, message.Text)
	for _, to := range message.To {
		m.AddRecipient(to)
	}
	_, id, err := mg.client.Send(m)
	if err != nil {
		return nil, err
	}

	return &MessageID{
		ID: id,
	}, nil
}
