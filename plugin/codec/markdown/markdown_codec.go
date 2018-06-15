package markdown

import (
	"sync"

	bluemonday "github.com/microcosm-cc/bluemonday"
	blackfriday "github.com/russross/blackfriday"
)

const (
	ContentTypeMARKDOWN   string = "application/markdown"
	FileExtensionMARKDOWN string = ".md"
)

var validMarkdownContentTypes = []string{
	"application/markdown",
	"text/markdown",
}

// MarkdownCodec converts objects to and from MARKDOWN.
type MarkdownCodec struct {
	config Config
	lock   *sync.RWMutex
}

// New returns a new MARKDOWN serializer
func New(cfg ...Config) *MarkdownCodec {
	c := DefaultConfig().Merge(cfg)
	return &MarkdownCodec{config: c}
}

// Converts an object to MARKDOWN.
func (c *MarkdownCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return markdownEncoding.Marshal(object)
}

// Unmarshal converts MARKDOWN into an object.
func (c *MarkdownCodec) Unmarshal(data []byte, obj interface{}) error {
	return markdownEncoding.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (c *MarkdownCodec) ContentType() string {
	return ContentTypeMARKDOWN
}

// FileExtension returns the file extension for this codec.
func (c *MarkdownCodec) FileExtension() string {
	return FileExtensionMARKDOWN
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *MarkdownCodec) CanMarshalWithCallback() bool {
	return false
}

func (c *MarkdownCodec) ContentTypeSupported(contentType string) bool {
	for _, supportedType := range validMarkdownContentTypes {
		if supportedType == contentType {
			return true
		}
	}
	return contentType == c.ContentType()
}
