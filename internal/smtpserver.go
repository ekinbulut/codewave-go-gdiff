package internal

import "net/smtp"

type SmtpServer struct {
	Host string
	Port string
	User string
	Pass string
}

func NewSmtpServer(host string, port string, user string, pass string) *SmtpServer {
	return &SmtpServer{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
	}
}

// return address
func (s *SmtpServer) GetAddress() string {
	return s.Host + ":" + s.Port
}

// send email
func (s *SmtpServer) SendEmail(from string, to string, subject string, body string) error {

	auth := smtp.PlainAuth("", s.User, s.Pass, s.Host)
	err := smtp.SendMail(s.GetAddress(), auth, from, []string{to}, []byte(body))
	if err != nil {
		return err
	}
	return nil
}
