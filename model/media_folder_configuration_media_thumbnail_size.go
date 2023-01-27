package model

type MediaFolderConfigurationMediaThumbnailSize struct {
	MediaFolderConfiguration   *MediaFolderConfiguration `json:"mediaFolderConfiguration,omitempty"`
	MediaFolderConfigurationId string                    `json:"mediaFolderConfigurationId"` // required
	MediaThumbnailSize         *MediaThumbnailSize       `json:"mediaThumbnailSize,omitempty"`
	MediaThumbnailSizeId       string                    `json:"mediaThumbnailSizeId"` // required
}
