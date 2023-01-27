package model

type NewsletterRecipientTag struct {
	NewsletterRecipient   *NewsletterRecipient `json:"newsletterRecipient,omitempty"`
	NewsletterRecipientId string               `json:"newsletterRecipientId"` // required
	Tag                   *Tag                 `json:"tag,omitempty"`
	TagId                 string               `json:"tagId"` // required
}
