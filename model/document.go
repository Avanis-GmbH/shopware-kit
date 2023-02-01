package model

import "time"

type Document struct {
	Config               interface{}   `json:"config,omitempty"` // required
	CreatedAt            *time.Time    `json:"createdAt,omitempty"`
	CustomFields         interface{}   `json:"customFields,omitempty"`
	DeepLinkCode         string        `json:"deepLinkCode,omitempty"` // required
	DependentDocuments   []Document    `json:"dependentDocuments,omitempty"`
	DocumentMediaFile    *Media        `json:"documentMediaFile,omitempty"`
	DocumentMediaFileId  string        `json:"documentMediaFileId,omitempty"`
	DocumentType         *DocumentType `json:"documentType,omitempty"`
	DocumentTypeId       string        `json:"documentTypeId,omitempty"` // required
	FileType             string        `json:"fileType,omitempty"`       // required
	Id                   string        `json:"id,omitempty"`
	Order                *Order        `json:"order,omitempty"`
	OrderId              string        `json:"orderId,omitempty"` // required
	OrderVersionId       string        `json:"orderVersionId,omitempty"`
	ReferencedDocument   *Document     `json:"referencedDocument,omitempty"`
	ReferencedDocumentId string        `json:"referencedDocumentId,omitempty"`
	Sent                 bool          `json:"sent,omitempty"`
	Static               bool          `json:"static,omitempty"`
	UpdatedAt            *time.Time    `json:"updatedAt,omitempty"`
}
