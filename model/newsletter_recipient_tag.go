package model

type NewsletterRecipientTag struct {
	NewsletterRecipient   *NewsletterRecipient `json:"newsletterRecipient,omitempty"`
	NewsletterRecipientId string               `json:"newsletterRecipientId,omitempty"`
	Tag                   *Tag                 `json:"tag,omitempty"`
	TagId                 string               `json:"tagId,omitempty"`
}

type NewsletterRecipientTagCollection struct {
	EntityCollection

	Data []NewsletterRecipientTag `json:"data"`
}
