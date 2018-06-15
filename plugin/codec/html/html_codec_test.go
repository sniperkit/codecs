package html

import (
	"testing"

	// external
	"github.com/stretchr/testify/assert"

	// internal
	codecs "github.com/sniperkit/codecs/pkg"
	constants "github.com/sniperkit/codecs/pkg/constants"
)

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(HtmlCodec))

}

func TestMarshal(t *testing.T) {

	codec := new(HtmlCodec)

	obj := make(map[string]string)
	obj["name"] = "Tyler"
	expectedResult := []byte{0x15, 0x0, 0x0, 0x0, 0x2, 0x6e, 0x61, 0x6d, 0x65, 0x0, 0x6, 0x0, 0x0, 0x0, 0x54, 0x79, 0x6c, 0x65, 0x72, 0x0, 0x0}

	htmlData, htmlError := codec.Marshal(obj, nil)

	if htmlError != nil {
		t.Errorf("Shouldn't return error: %s", htmlError)
	}

	assert.Equal(t, htmlData, expectedResult)

}

func TestUnmarshal(t *testing.T) {

	codec := new(HtmlCodec)
	htmlData := []byte{0x15, 0x0, 0x0, 0x0, 0x2, 0x6e, 0x61, 0x6d, 0x65, 0x0, 0x6, 0x0, 0x0, 0x0, 0x54, 0x79, 0x6c, 0x65, 0x72, 0x0, 0x0}
	var object map[string]interface{}

	err := codec.Unmarshal(htmlData, &object)

	if assert.Nil(t, err) {
		assert.Equal(t, "Tyler", object["name"])
	}

}

func TestResponseContentType(t *testing.T) {

	codec := new(HtmlCodec)
	assert.Equal(t, codec.ContentType(), ContentTypeHTML)
}

func TestFileExtension(t *testing.T) {

	codec := new(HtmlCodec)
	assert.Equal(t, FileExtensionHTML, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	codec := new(HtmlCodec)
	assert.False(t, codec.CanMarshalWithCallback())

}
