package email

type Sender interface {
	Send(*Message) (*MessageID, error)
}
