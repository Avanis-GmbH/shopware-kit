package model

type MediaFolderConfigurationMediaThumbnailSize struct {
	MediaFolderConfiguration   *MediaFolderConfiguration `json:"mediaFolderConfiguration,omitempty"`
	MediaFolderConfigurationId string                    `json:"mediaFolderConfigurationId,omitempty"`
	MediaThumbnailSize         *MediaThumbnailSize       `json:"mediaThumbnailSize,omitempty"`
	MediaThumbnailSizeId       string                    `json:"mediaThumbnailSizeId,omitempty"`
}

type MediaFolderConfigurationMediaThumbnailSizeCollection struct {
	EntityCollection

	Data []MediaFolderConfigurationMediaThumbnailSize `json:"data"`
}
