package mail

import (
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"fmt"
	"github.com/koin-bet/koin.backend/pkg/util"
	"os"
)

func SendMail(hash string, email string) {
	d := gomail.NewPlainDialer(
		os.Getenv("mail_host"),
		util.GetEnvOrDefaultInt("mail_port", 0),
		os.Getenv("mail"),
		os.Getenv("mail_pwd"),
	)
	d.TLSConfig = &tls.Config{}
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("mail"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Koin.bet - Your hash is here")
	m.SetBody("text/html", "Hello, to join the game with your koins, set your hash with this: <br> "+hash)
	fmt.Println("Email sent to " + email)
	fmt.Println("Config is  ", d)
	e := d.DialAndSend(m)
	if e != nil {
		fmt.Errorf("ERROR MAIL: " + e.Error() + "\n")
	}
}
