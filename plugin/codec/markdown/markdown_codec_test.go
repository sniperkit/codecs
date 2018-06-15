package markdown

import (
	"testing"

	// external
	assert "github.com/stretchr/testify/assert"

	// internal
	codecs "github.com/sniperkit/codecs/pkg"
	constants "github.com/sniperkit/codecs/pkg/constants"
)

var codec MarkdownCodec

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(MarkdownCodec), "MarkdownCodec")

}

func TestMarshal(t *testing.T) {

	obj := make(map[string]string)
	obj["name"] = "Mat"

	markdownString, markdownError := codec.Marshal(obj, nil)

	if markdownError != nil {
		t.Errorf("Shouldn't return error: %s", markdownError)
	}

	assert.Equal(t, string(markdownString), `name: Mat`)

}

func TestUnmarshal(t *testing.T) {

	markdownString := `name: Mat`
	var object map[string]interface{}

	err := codec.Unmarshal([]byte(markdownString), &object)

	if err != nil {
		t.Errorf("Shouldn't return error: %s", err)
	}

	assert.Equal(t, "Mat", object["name"])

}

func TestResponseContentType(t *testing.T) {

	assert.Equal(t, codec.ContentType(), ContentTypeMARKDOWN)

}

func TestFileExtension(t *testing.T) {

	assert.Equal(t, FileExtensionMARKDOWN, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	assert.False(t, codec.CanMarshalWithCallback())

}
