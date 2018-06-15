package rss

import (
	"sync"

	rssEncoding "github.com/sniperkit/gofeed/pkg"
)

const (
	ContentType   string = "application/rss"
	FileExtension string = ".rss"
)

var validRssContentTypes = []string{
	"application/rss",
	"text/rss",
}

// RssCodec converts objects to and from RSS.
type RssCodec struct {
	config Config
	lock   *sync.RWMutex
}

// New returns a new RSS serializer
func New(cfg ...Config) *RssCodec {
	c := DefaultConfig().Merge(cfg)
	return &RssCodec{config: c}
}

// Converts an object to RSS.
func (c *RssCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return rssEncoding.Marshal(object)
}

// Unmarshal converts RSS into an object.
func (c *RssCodec) Unmarshal(data []byte, obj interface{}) error {
	return rssEncoding.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (c *RssCodec) ContentType() string {
	return ContentType
}

// FileExtension returns the file extension for this codec.
func (c *RssCodec) FileExtension() string {
	return FileExtension
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *RssCodec) CanMarshalWithCallback() bool {
	return false
}

func (c *RssCodec) ContentTypeSupported(contentType string) bool {
	for _, supportedType := range validRssContentTypes {
		if supportedType == contentType {
			return true
		}
	}
	return contentType == c.ContentType()
}
