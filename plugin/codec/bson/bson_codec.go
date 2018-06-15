package bson

import (
	"sync"

	// external
	"gopkg.in/mgo.v2/bson"
)

const (
	ContentType   string = "application/bson"
	FileExtension string = ".bson"
)

// BsonCodec converts objects to and from BSON.
type BsonCodec struct {
	config Config
	lock   *sync.RWMutex
}

// New returns a new BSON serializer
func New(cfg ...Config) *BsonCodec {
	c := DefaultConfig().Merge(cfg)
	return &BsonCodec{config: c}
}

// Marshal converts an object to BSON.
func (b *BsonCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return bson.Marshal(object)
}

// Unmarshal converts JSON into an object.
func (b *BsonCodec) Unmarshal(data []byte, obj interface{}) error {
	return bson.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (b *BsonCodec) ContentType() string {
	return ContentType
}

// FileExtension returns the file extension for this codec.
func (b *BsonCodec) FileExtension() string {
	return FileExtension
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (b *BsonCodec) CanMarshalWithCallback() bool {
	return false
}
