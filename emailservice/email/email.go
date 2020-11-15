package email

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

var (
	host     string
	port     string
	sender   string
	password string
)

//SendEmail to the user
func SendEmail(to string, token string) (string, string, error) {
	body, err := CreateVerificationEmail(token)
	if err != nil {
		return "", "", err
	}

	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Create a new template

	err = mg.CreateTemplate(ctx, &mailgun.Template{
		Name: "Verification-Email",
		Version: mailgun.TemplateVersion{
			Template: body,
			Engine:   mailgun.TemplateEngineGo,
			Tag:      "v1",
		},
	})

	// Give time for template to show up in the system.
	time.Sleep(time.Second * 1)

	// Create a new message with template
	m := mg.NewMessage("Syahrul mailgun@mail.oncecard.com", "Template example", "")
	m.SetTemplate("Verification-Email")

	// Add recipients
	m.AddRecipient(to)

	resp, id, err := mg.Send(ctx, m)
	return resp, id, err
}

//SendSimpleMessage message sends the simple message to the user
func SendSimpleMessage(recipient string) (string, error) {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		fmt.Sprintf("Syahrul %s", "mailgun@mail.oncecard.com"),
		"Hello",
		"Testing some Mailgun awesomeness!",
		recipient,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
