package jsonp

import (
	"testing"

	// external
	assert "github.com/stretchr/testify/assert"

	// internal
	codecs "github.com/sniperkit/codecs/pkg"
	constants "github.com/sniperkit/codecs/pkg/constants"
)

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(JsonPCodec), "JsonPCodec")

}

func TestResponseContentType(t *testing.T) {

	codec := new(JsonPCodec)
	assert.Equal(t, codec.ContentType(), ContentTypeJSONP)

}

func TestFileExtension(t *testing.T) {

	codec := new(JsonPCodec)
	assert.Equal(t, FileExtensionJSONP, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	codec := new(JsonPCodec)
	assert.True(t, codec.CanMarshalWithCallback())

}

func TestMarshal(t *testing.T) {

	codec := new(JsonPCodec)

	obj := make(map[string]string)
	obj["name"] = "Mat"

	jsonPString, jsonPError := codec.Marshal(obj, map[string]interface{}{constants.OptionKeyClientCallback: "candyCorn", "not-relevant": true})

	if jsonPError != nil {
		t.Errorf("Shouldn't return error: %s", jsonPError)
	}

	assert.Equal(t, string(jsonPString), `candyCorn({"name":"Mat"});`)

}

func TestMarshal_WithContext(t *testing.T) {

	codec := new(JsonPCodec)

	obj := make(map[string]string)
	obj["name"] = "Mat"

	jsonPString, jsonPError := codec.Marshal(obj, map[string]interface{}{constants.OptionKeyClientCallback: "candyCorn", OptionKeyClientContext: "halloween", "not-relevant": true})

	if jsonPError != nil {
		t.Errorf("Shouldn't return error: %s", jsonPError)
	}

	assert.Equal(t, string(jsonPString), `candyCorn({"name":"Mat"},"halloween");`)

}

func TestMarshal_WithoutCallback(t *testing.T) {

	codec := new(JsonPCodec)

	obj := make(map[string]string)
	obj["name"] = "Mat"

	_, jsonPError := codec.Marshal(obj, nil)

	assert.Equal(t, jsonPError, ErrorMissingCallback)
}

func TestUnmarshal(t *testing.T) {

	codec := new(JsonPCodec)

	jsonString := `{"name":"Mat"}`
	var object map[string]interface{}

	jsonPError := codec.Unmarshal([]byte(jsonString), &object)

	assert.Equal(t, jsonPError, ErrorUnmarshalNotSupported)
}
