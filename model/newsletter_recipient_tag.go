package model

type NewsletterRecipientTag struct {
	NewsletterRecipient   *NewsletterRecipient `json:"newsletterRecipient,omitempty"`
	NewsletterRecipientId string               `json:"newsletterRecipientId,omitempty"` // required
	Tag                   *Tag                 `json:"tag,omitempty"`
	TagId                 string               `json:"tagId,omitempty"` // required
}
