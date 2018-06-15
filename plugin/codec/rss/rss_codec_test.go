package rss

import (
	"testing"

	// external
	assert "github.com/stretchr/testify/assert"

	// internal
	codecs "github.com/sniperkit/codecs/pkg"
	constants "github.com/sniperkit/codecs/pkg/constants"
)

var codec RssCodec

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(RssCodec), "RssCodec")

}

func TestMarshal(t *testing.T) {

	obj := make(map[string]string)
	obj["name"] = "Mat"

	rssString, rssError := codec.Marshal(obj, nil)

	if rssError != nil {
		t.Errorf("Shouldn't return error: %s", rssError)
	}

	assert.Equal(t, string(rssString), `name: Mat`)

}

func TestUnmarshal(t *testing.T) {

	rssString := `name: Mat`
	var object map[string]interface{}

	err := codec.Unmarshal([]byte(rssString), &object)

	if err != nil {
		t.Errorf("Shouldn't return error: %s", err)
	}

	assert.Equal(t, "Mat", object["name"])

}

func TestResponseContentType(t *testing.T) {

	assert.Equal(t, codec.ContentType(), ContentTypeRSS)

}

func TestFileExtension(t *testing.T) {

	assert.Equal(t, FileExtensionRSS, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	assert.False(t, codec.CanMarshalWithCallback())

}
