package brevo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	brevo "github.com/getbrevo/brevo-go/lib"
	"keyz/backend/prisma/db"
)

type emailBody struct {
	Sender *brevo.SendSmtpEmailSender `json:"sender,omitempty"`
	// Mandatory if messageVersions are not passed, ignored if messageVersions are passed. List of email addresses and names (optional) of the recipients. For example, [{\"name\":\"Jimmy\", \"email\":\"jimmy98@example.com\"}, {\"name\":\"Joe\", \"email\":\"joe@example.com\"}]
	To []brevo.SendSmtpEmailTo `json:"to,omitempty"`
	// List of email addresses and names (optional) of the recipients in bcc
	Bcc []brevo.SendSmtpEmailBcc `json:"bcc,omitempty"`
	// List of email addresses and names (optional) of the recipients in cc
	Cc      []brevo.SendSmtpEmailCc     `json:"cc,omitempty"`
	ReplyTo *brevo.SendSmtpEmailReplyTo `json:"replyTo,omitempty"`
	// Subject of the message. Mandatory if 'templateId' is not passed
	Subject string `json:"subject,omitempty"`
	// Pass the absolute URL (no local file) or the base64 content of the attachment along with the attachment name (Mandatory if attachment content is passed). For example, `[{\"url\":\"https://attachment.domain.com/myAttachmentFromUrl.jpg\", \"name\":\"myAttachmentFromUrl.jpg\"}, {\"content\":\"base64 example content\", \"name\":\"myAttachmentFromBase64.jpg\"}]`. Allowed extensions for attachment file: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub, eps, odt, mp3, m4a, m4v, wma, ogg, flac, wav, aif, aifc, aiff, mp4, mov, avi, mkv, mpeg, mpg, wmv, pkpass and xlsm ( If 'templateId' is passed and is in New Template Language format then both attachment url and content are accepted. If template is in Old template Language format, then 'attachment' is ignored )
	Attachment []brevo.SendSmtpEmailAttachment `json:"attachment,omitempty"`
	// Id of the template.
	TemplateId int64 `json:"templateId,omitempty"`
	// Pass the set of attributes to customize the template. For example, {\"FNAME\":\"Joe\", \"LNAME\":\"Doe\"}. It's considered only if template is in New Template Language format.
	Params map[string]any `json:"params,omitempty"`
}

func buildBody(fromName string, toEmail string, cc []string, replyTo string, templateId int64, subject string, params map[string]any) emailBody {
	body := emailBody{
		Sender: &brevo.SendSmtpEmailSender{
			Name:  fromName,
			Email: "contact@keyz-app.fr",
		},
		To: []brevo.SendSmtpEmailTo{
			{
				Email: toEmail,
			},
		},
		Subject:    subject,
		TemplateId: templateId,
		Params:     params,
	}
	if len(cc) > 0 {
		body.Cc = []brevo.SendSmtpEmailCc{}
		for _, email := range cc {
			body.Cc = append(body.Cc, brevo.SendSmtpEmailCc{Email: email})
		}
	}
	if len(replyTo) > 0 {
		body.ReplyTo = &brevo.SendSmtpEmailReplyTo{Email: replyTo}
	}
	return body
}

func callBrevo(fromName string, toEmail string, cc []string, replyTo string, templateId int64, subject string, params map[string]any) (string, error) {
	apiURL := "https://api.brevo.com/v3/smtp/email"
	apiKey := os.Getenv("BREVO_API_KEY")

	body := buildBody(fromName, toEmail, cc, replyTo, templateId, subject, params)
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func SendEmailInvite(invite db.LeaseInviteModel, userExists bool) (string, error) {
	ownerName := invite.Property().Owner().Name()
	ownerEmail := invite.Property().Owner().Email
	var inviteLink string
	if userExists {
		inviteLink = os.Getenv("WEB_PUBLIC_URL") + "/login/invite/" + invite.ID
	} else {
		inviteLink = os.Getenv("WEB_PUBLIC_URL") + "/register/invite/" + invite.ID
	}
	params := map[string]any{
		"ownerName":  ownerName,
		"inviteLink": inviteLink,
	}
	subject := "You've been invited to join a property on Keyz"

	return callBrevo(ownerName+" via Keyz", invite.TenantEmail, []string{}, ownerEmail, 1, subject, params)
}

func SendNewDamage(lease db.LeaseModel) (string, error) {
	tenantName := lease.Tenant().Name()
	tenantEmail := lease.Tenant().Email
	propertyName := lease.Property().Name
	params := map[string]any{
		"tenantName":   tenantName,
		"propertyName": propertyName,
		"damageLink":   os.Getenv("WEB_PUBLIC_URL") + "/real-property/details/" + lease.PropertyID,
	}
	subject := "A new damage has been created in " + propertyName

	return callBrevo(tenantName+" via Keyz", lease.Property().Owner().Email, []string{}, tenantEmail, 5, subject, params)
}

func SendNewContactMessage(cm db.ContactMessageModel) (string, error) {
	name := cm.Firstname + " " + cm.Lastname
	params := map[string]any{
		"id":          cm.ID,
		"senderName":  name,
		"senderEmail": cm.Email,
		"subject":     cm.Subject,
		"message":     cm.Message,
	}
	subject := "[New message from keyz] " + cm.Subject

	return callBrevo(name+" via keyz-app.fr", "contact@keyz-app.fr", []string{}, cm.Email, 4, subject, params)
}
