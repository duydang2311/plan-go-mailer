package handlers

import (
	"log"

	"plan/internal/runtime"

	"github.com/goccy/go-json"
	"github.com/nats-io/nats.go"
	"github.com/resend/resend-go/v2"
)

type Message struct {
	From    string
	To      []string
	Subject string
	Text    string
	Html    string
}

func SendMailHandler(runtime runtime.Runtime) (*nats.Subscription, error) {
	return runtime.Nats.Subscribe("mails.send", func(msg *nats.Msg) {
		var m Message
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Println(err)
			return
		}

		sent, err := runtime.Mailer.Emails.Send(&resend.SendEmailRequest{
			From:    m.From,
			To:      m.To,
			Subject: m.Subject,
			Html:    m.Html,
			Text:    m.Text,
		})

		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("sent mail %s", sent.Id)
	})
}
