package email

import (
	"strconv"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 4002
const CONFIG_EMAIL = "adrestiapersephone@gmail.com"
const CONFIG_PASSWORD = "adrestiapersephone123"

func sendEmail(destination []string, subject string, message string) (error error){

	//body := "From: " + CONFIG_EMAIL + "\n" +
	//	"To: " + strings.Join(destination, ",") + "\n" +
	//	"Subject: " + subject + "\n\n" +
	//	message
	//
	//auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	//smtpAddr := fmt.Sprintf("")

	return nil
}


func OpenEmail(destination []string, code int) (error error) {

	subject := "Karcis Company"
	message := "Enter your verification code: " + strconv.Itoa(code)

	err := sendEmail(destination, subject, message)

	return err
}
