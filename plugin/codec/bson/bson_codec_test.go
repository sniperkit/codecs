package bson

import (
	"testing"

	// external
	"github.com/stretchr/testify/assert"

	// internal
	codecs "github.com/sniperkit/codecs/pkg"
	constants "github.com/sniperkit/codecs/pkg/constants"
)

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(BsonCodec))

}

func TestMarshal(t *testing.T) {

	codec := new(BsonCodec)

	obj := make(map[string]string)
	obj["name"] = "Tyler"
	expectedResult := []byte{0x15, 0x0, 0x0, 0x0, 0x2, 0x6e, 0x61, 0x6d, 0x65, 0x0, 0x6, 0x0, 0x0, 0x0, 0x54, 0x79, 0x6c, 0x65, 0x72, 0x0, 0x0}

	bsonData, bsonError := codec.Marshal(obj, nil)

	if bsonError != nil {
		t.Errorf("Shouldn't return error: %s", bsonError)
	}

	assert.Equal(t, bsonData, expectedResult)

}

func TestUnmarshal(t *testing.T) {

	codec := new(BsonCodec)
	bsonData := []byte{0x15, 0x0, 0x0, 0x0, 0x2, 0x6e, 0x61, 0x6d, 0x65, 0x0, 0x6, 0x0, 0x0, 0x0, 0x54, 0x79, 0x6c, 0x65, 0x72, 0x0, 0x0}
	var object map[string]interface{}

	err := codec.Unmarshal(bsonData, &object)

	if assert.Nil(t, err) {
		assert.Equal(t, "Tyler", object["name"])
	}

}

func TestResponseContentType(t *testing.T) {

	codec := new(BsonCodec)
	assert.Equal(t, codec.ContentType(), ContentTypeBSON)
}

func TestFileExtension(t *testing.T) {

	codec := new(BsonCodec)
	assert.Equal(t, FileExtensionBSON, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	codec := new(BsonCodec)
	assert.False(t, codec.CanMarshalWithCallback())

}
