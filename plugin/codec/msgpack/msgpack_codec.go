package msgpack

import (
	"bytes"
	"sync"

	// external
	codec "github.com/ugorji/go/codec"
)

// MsgpackCodec converts objects to and from Msgpack.
type MsgpackCodec struct {
	config Config
	lock   *sync.RWMutex
}

const (
	ContentType   string = "application/x-msgpack"
	FileExtension string = ".msgpack"
)

var msgpackHandle codec.MsgpackHandle

// New returns a new Msgpack serializer
func New(cfg ...Config) *MsgpackCodec {
	c := DefaultConfig().Merge(cfg)
	return &MsgpackCodec{config: c}
}

// Converts an object to Msgpack.
func (c *MsgpackCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {

	byteBuffer := new(bytes.Buffer)
	enc := codec.NewEncoder(byteBuffer, &msgpackHandle)
	encErr := enc.Encode(object)

	return byteBuffer.Bytes(), encErr
}

// Unmarshal converts Msgpack into an object.
func (c *MsgpackCodec) Unmarshal(data []byte, obj interface{}) error {

	dec := codec.NewDecoder(bytes.NewReader(data), &msgpackHandle)
	return dec.Decode(&obj)
}

// ContentType returns the content type for this codec.
func (c *MsgpackCodec) ContentType() string {
	return ContentType
}

// FileExtension returns the file extension for this codec.
func (c *MsgpackCodec) FileExtension() string {
	return FileExtension
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *MsgpackCodec) CanMarshalWithCallback() bool {
	return false
}
