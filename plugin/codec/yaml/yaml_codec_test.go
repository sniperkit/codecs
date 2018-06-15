package yaml

import (
	"testing"

	// external
	assert "github.com/stretchr/testify/assert"

	// internal
	codecs "github.com/sniperkit/codecs/pkg"
	constants "github.com/sniperkit/codecs/pkg/constants"
)

var codec YamlCodec

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(YamlCodec), "YamlCodec")

}

func TestMarshal(t *testing.T) {

	obj := make(map[string]string)
	obj["name"] = "Mat"

	yamlString, yamlError := codec.Marshal(obj, nil)

	if yamlError != nil {
		t.Errorf("Shouldn't return error: %s", yamlError)
	}

	assert.Equal(t, string(yamlString), `name: Mat`)

}

func TestUnmarshal(t *testing.T) {

	yamlString := `name: Mat`
	var object map[string]interface{}

	err := codec.Unmarshal([]byte(yamlString), &object)

	if err != nil {
		t.Errorf("Shouldn't return error: %s", err)
	}

	assert.Equal(t, "Mat", object["name"])

}

func TestResponseContentType(t *testing.T) {

	assert.Equal(t, codec.ContentType(), ContentTypeYAML)

}

func TestFileExtension(t *testing.T) {

	assert.Equal(t, FileExtensionYAML, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	assert.False(t, codec.CanMarshalWithCallback())

}
