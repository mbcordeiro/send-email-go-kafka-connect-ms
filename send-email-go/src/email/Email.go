package email

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type Email struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func NewEmail() *Email {
	return &Email{}
}

type MailSender struct {
	From   string
	Dailer *gomail.Dialer
}

func NewMailSender() *MailSender {
	return &MailSender{}
}

func (mailsender *MailSender) Send(emailChan chan Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mailsender.From)
	for ec := range emailChan {
		m.SetHeader("Subject", ec.Subject)
		m.SetBody("text/html", ec.Body)
		for _, to := range ec.Emails {
			m.SetHeader("To", to)
			if err := mailsender.Dailer.DialAndSend(m); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
