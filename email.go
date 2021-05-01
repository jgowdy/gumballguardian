
type EmailSender interface {
	// TODO: Some interface for polling via IMAP or POP3 or ???
	SendPlainTextEmail(targetEmail string, subject string, body string)
}

type EmailReceiver interface {
	// TODO: Some interface for polling via IMAP or POP3

}