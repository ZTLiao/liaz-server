package utils

import "gopkg.in/gomail.v2"

func SendMail(username string, password string, host string, port int, nickname string, toUsername string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(username, nickname))
	m.SetHeader("To", toUsername)
	m.SetHeader("Subject", subject)
	m.SetHeader("text/html", body)
	d := gomail.NewDialer(host, port, username, password)
	return d.DialAndSend(m)
}
