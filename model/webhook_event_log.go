package model

import "time"

type WebhookEventLog struct {
	AppName                  string      `json:"appName,omitempty"`
	AppVersion               string      `json:"appVersion,omitempty"`
	CreatedAt                time.Time   `json:"createdAt,omitempty"`
	CustomFields             interface{} `json:"customFields,omitempty"`
	DeliveryStatus           string      `json:"deliveryStatus"` // required
	EventName                string      `json:"eventName"`      // required
	Id                       string      `json:"id,omitempty"`
	ProcessingTime           float64     `json:"processingTime"`
	RequestContent           interface{} `json:"requestContent,omitempty"`
	ResponseContent          interface{} `json:"responseContent,omitempty"`
	ResponseReasonPhrase     string      `json:"responseReasonPhrase,omitempty"`
	ResponseStatusCode       float64     `json:"responseStatusCode"`
	SerializedWebhookMessage interface{} `json:"serializedWebhookMessage,omitempty"`
	Timestamp                float64     `json:"timestamp"`
	UpdatedAt                time.Time   `json:"updatedAt,omitempty"`
	Url                      string      `json:"url"`         // required
	WebhookName              string      `json:"webhookName"` // required
}
