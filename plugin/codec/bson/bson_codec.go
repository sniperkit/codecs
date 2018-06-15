package bson

import (
	// external
	"gopkg.in/mgo.v2/bson"
	// internal
	// constants "github.com/sniperkit/codecs/pkg/constants"
)

const (
	ContentTypeBSON   string = "application/bson"
	FileExtensionBSON string = ".bson"
)

// BsonCodec converts objects to and from BSON.
type BsonCodec struct{}

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
	return ContentTypeBSON
}

// FileExtension returns the file extension for this codec.
func (b *BsonCodec) FileExtension() string {
	return FileExtensionBSON
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (b *BsonCodec) CanMarshalWithCallback() bool {
	return false
}
