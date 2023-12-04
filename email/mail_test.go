package email

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	sender := "ericwangcatch@gmail.com"
	em := EMail{
		Sender:   sender,
		Receiver: "palefever98@gmail.com",
		Subject:  "Go mail test",
		HTMLBody: "Go mail test, don't reply",
	}
	err := SendMail(em)
	if err != nil {
		t.Fatal(err)
	}
}
