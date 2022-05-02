package srv

type Notification struct {
	Message    string
	SmtpServer *SmtpServer
}

func NewNotification(message string, smtpServer *SmtpServer) *Notification {
	return &Notification{
		Message:    message,
		SmtpServer: NewSmtpServer("smtp.google.com", 465, "", ""),
	}
}

// send email
func (n *Notification) SendEmail() {
	n.SmtpServer.SendEmail("", "", "", n.Message)
}
