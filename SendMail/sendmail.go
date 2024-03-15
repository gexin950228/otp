package SendMail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"otp/logSource"
)

type SendMail struct {
	Code int
	Msg  string
}

func SendEmail(userId int, Subject string, body string, to string) (status SendMail) {
	msg := gomail.NewMessage()
	mailMsg := LoadMailConfig("../conf/mail.json")
	msg.SetHeader("Subject", Subject)
	msg.SetHeader("From", mailMsg.From)
	msg.SetHeader("To", to)
	msg.SetBody("text/html", body)
	n := gomail.NewDialer(mailMsg.Server, mailMsg.Port, mailMsg.User, mailMsg.Password)
	if err := n.DialAndSend(msg); err != nil {
		logSource.Log.Error(err.Error())
		status.Code = 2
		status.Msg = err.Error()
	} else {
		status.Code = 1
		status.Msg = fmt.Sprintf("发送给%s的邮件发送成功。", to)
		logSource.Log.Info("邮件发送成功")
	}
	return status
}
