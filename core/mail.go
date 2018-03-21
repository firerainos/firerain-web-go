package core

import (
	"net/smtp"
)

func SendMail(toEmail string) error {
	username := Conf.Smtp.Username
	auth := smtp.PlainAuth("", username, Conf.Smtp.Password, Conf.Smtp.Host)
	subject := "FireRainOS内测申请审核通过"
	body := "您已通过FireRainOS内测申请审核，请及时(过时将关闭进群审核)加入qq群:615676312 (入群请填写申请时用的邮箱)来进一步获取内部内测消息及问题建议反馈"

	msg := []byte("To: " + toEmail + "\nFrom: " + username + "\nSubject: " + subject + "\n\n" + body)
	return smtp.SendMail(Conf.Smtp.Host+":25", auth, Conf.Smtp.Username, []string{toEmail}, msg)
}
