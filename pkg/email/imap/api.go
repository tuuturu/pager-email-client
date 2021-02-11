package imap

import (
	"fmt"
	"net/url"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/tuuturu/pager-email-client/pkg/core"
)

func RetrieveEmails(serverURL *url.URL, username, password string) (emails []core.Email, err error) {
	c, err := client.DialTLS(serverURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("opening IMAP client: %w", err)
	}

	defer func() { _ = c.Close() }()

	err = c.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("logging into IMAP server: %w", err)
	}

	mbox, err := c.Select("INBOX", true)
	if err != nil {
		return nil, fmt.Errorf("selecting inbox: %w", err)
	}

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		from = mbox.Messages - 3
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)

	err = c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	if err != nil {
		return nil, err
	}

	emails = make([]core.Email, 0)
	for msg := range messages {
		email := core.Email{
			Subject: msg.Envelope.Subject,
		}

		if len(msg.Envelope.To) > 0 {
			email.To = msg.Envelope.To[0].Address()
		}

		if len(msg.Envelope.From) > 0 {
			email.From = msg.Envelope.From[0].Address()
		}

		emails = append(emails, email)
	}

	return emails, nil
}
