package srv

import "net/smtp"

type SmtpServer struct {
	Host string
	Port int
	User string
	Pass string
}

func NewSmtpServer(host string, port int, user string, pass string) *SmtpServer {
	return &SmtpServer{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
	}
}

// send email
func (s *SmtpServer) SendEmail(from string, to string, subject string, body string) error {

	err := smtp.SendMail(s.Host+":"+string(s.Port), nil, from, []string{to}, []byte(body))
	if err != nil {
		return err
	}
	return nil
}
