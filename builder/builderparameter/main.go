package main

import (
	"fmt"
	"strings"
)

type email struct {
	To      []string
	From    string
	Subject string
	Body    string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.To = append(b.email.To, to)
	return b
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("Email should contain @")
	}
	b.email.From = from
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.Subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.Body = body
	return b
}

func sendMailImpl(email *email) {

}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

// main is the entry point of the application. It demonstrates the usage of the
// builder pattern to construct an email with specified parameters such as
// sender, recipient, subject, and body. The constructed email is then sent
// using the SendEmail function.
func main() {
	fmt.Println("---------Builder Paramter----------")
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})

	fmt.Println("Email sent successfully")

}
