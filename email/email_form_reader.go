package email

import (
	"embed"
	"log"
)

//go:embed verify_email_form.html
var formFile embed.FS

func MustReadEmailForm() string {
	data, err := formFile.ReadFile("verify_email_form.html")
	if err != nil {
		log.Println(err)
	}
	return string(data)
}
