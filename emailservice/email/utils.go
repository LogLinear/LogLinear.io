package email

import (
	"bytes"
	"html/template"
)

//CreateVerificationEmail will parse the template.html file and include the token inside
//Use this to get a stringified body of the email
func CreateVerificationEmail(token string) (string, error) {
	templateData := struct {
		VerificationToken string
	}{
		VerificationToken: token,
	}
	// subject := "Verify your email @ OnceCard \n"
	t, err := template.ParseFiles("template.html")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		return "", err
	}

	body := buf.String()
	if err != nil {
		return "", err
	}
	return body, nil
}
