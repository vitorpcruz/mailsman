package mailer

import (
	"fmt"
	"log"

	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailsman/internal/domain"
)

type EmailJetMailer struct{}

func (mailer *EmailJetMailer) SendEmail(batch domain.EmailBatch) error {
	mailjetClient := mailjet.NewMailjetClient(
		"",
		"",
	)

	recipients := configureMailjetRecipients(batch)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "vpcgotestemail@gmail.com",
				Name:  "Me",
			},
			To:       recipients,
			Subject:  batch.Title.Value,
			HTMLPart: batch.Message.Value,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data: %+v\n", res)

	return nil
}

func configureMailjetRecipients(batch domain.EmailBatch) *mailjet.RecipientsV31 {
	list := mailjet.RecipientsV31{}

	for _, email := range batch.Recipients.EmailAddresses {
		list = append(
			list,
			mailjet.RecipientV31{Email: email.Value},
		)
	}

	return &list
}
