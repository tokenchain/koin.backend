package mail

import (
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"fmt"
	"os"
)

func SendMail(hash string, email string) {
	d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("email"), os.Getenv("emailpwd"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", "koinkoin.io@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Koinkoin.io - Your hash is here")
	m.SetBody("text/html", "Hello, to join the game with your koins, set your hash with this: <br> " + hash)
	e := d.DialAndSend(m)
	if e != nil {
		fmt.Errorf("ERROR MAIL: " + e.Error() + "\n")
	}
}