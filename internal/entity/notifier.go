package entity

type (
	Notifier struct {
		To         []string
		Subject    string
		Body       string
		Attachment File
	}
)
