package send

import (
	"io"
	"log"
	"text/template"

	gomail "gopkg.in/mail.v2"
)

type User struct {
	TaskName string
	Name     string
	Surname  string
	Email    string
	Birthday string
	Link     string
}

func Send(user User) {
	d := gomail.NewDialer("smtp.mail.ru", 465, "darhan.gruz@mail.ru", "Kvwi4MddUkfbLJu63tPT")
	d.StartTLSPolicy = gomail.MandatoryStartTLS
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles("index.html")

	m := gomail.NewMessage()
	m.SetHeader("From", "darhan.gruz@mail.ru")
	m.SetAddressHeader("To", user.Email, user.Name)
	m.SetHeader("Subject", "Тема сообщении")
	m.SetBodyWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, user)
	})

	if err := gomail.Send(s, m); err != nil {
		log.Printf("Could not send email to %q: %v", user.Email, err)
	}
	m.Reset()
	log.Printf("Сообщение на адрес %s отправлен!", user.Email)
}