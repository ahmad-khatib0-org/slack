package model

const (
	PostEmbedImage             PostEmbedType = "image"
	PostEmbedMessageAttachment PostEmbedType = "message_attachment"
	PostEmbedOpengraph         PostEmbedType = "opengraph"
	PostEmbedLink              PostEmbedType = "link"
	PostEmbedPermalink         PostEmbedType = "permalink"
	PostEmbedBoards            PostEmbedType = "boards"
)

type PostEmbedType string

type PostEmbed struct {
	Type PostEmbedType `json:"type"`

	// The URL of the embedded content. Used for image and OpenGraph embeds.
	URL string `json:"url,omitempty"`

	// Any additional data for the embedded content. Only used for OpenGraph embeds.
	Data any `json:"data,omitempty"`
}

func (pe *PostEmbed) Auditable() map[string]any {
	// filter out embedded content.
	return map[string]any{
		"type": pe.Type,
		"url":  pe.URL,
	}
}
