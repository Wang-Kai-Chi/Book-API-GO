package email

import (
	"embed"

	emailverifier "github.com/AfterShip/email-verifier"
	mail "gopkg.in/gomail.v2"
)

type EMail struct {
	Sender   string
	Receiver string
	Subject  string
	HTMLBody string
}

//go:embed password.txt
var embedKey embed.FS

func mustGetKey() []byte {
	key, err := embedKey.ReadFile("password.txt")
	if err != nil {
		panic(err)
	}
	return key
}

func SendMail(em EMail) error {
	password := string(mustGetKey())

	m := mail.NewMessage()
	m.SetHeader("From", em.Sender)

	m.SetHeader("To", em.Receiver)

	m.SetHeader("Subject", em.Subject)

	m.SetBody("text/html", em.HTMLBody)

	d := mail.NewDialer("smtp.gmail.com", 587, em.Sender, password)

	err := d.DialAndSend(m)
	return err
}

func VerifyEmail(email string) error {
	verifier := emailverifier.NewVerifier()

	ret, err := verifier.Verify(email)
	if !ret.Syntax.Valid {
		panic("email address syntax is invalid")
	}
	return err
}
