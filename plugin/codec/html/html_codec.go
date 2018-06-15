package html

import (
	"sync"

	// external
	"gopkg.in/mgo.v2/html"
)

const (
	ContentType   string = "text/html"
	FileExtension string = ".html"
)

// HtmlCodec converts objects to and from HTML.
type HtmlCodec struct {
	config Config
	lock   *sync.RWMutex
}

// New returns a new HTML serializer
func New(cfg ...Config) *HtmlCodec {
	c := DefaultConfig().Merge(cfg)
	return &HtmlCodec{config: c}
}

// Marshal converts an object to HTML.
func (b *HtmlCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return html.Marshal(object)
}

// Unmarshal converts JSON into an object.
func (b *HtmlCodec) Unmarshal(data []byte, obj interface{}) error {
	return html.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (b *HtmlCodec) ContentType() string {
	return ContentType
}

// FileExtension returns the file extension for this codec.
func (b *HtmlCodec) FileExtension() string {
	return FileExtension
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (b *HtmlCodec) CanMarshalWithCallback() bool {
	return false
}
