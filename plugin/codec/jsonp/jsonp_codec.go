package jsonp

import (
	"errors"
	"sync"

	jsonEncoding "encoding/json"

	// external
	stewstrings "github.com/stretchr/stew/strings"

	// internal
	constants "github.com/sniperkit/codecs/pkg/constants"
)

const (
	ContentTypeJSONP   string = "text/javascript"
	FileExtensionJSONP string = ".js"
)

var validJsonpContentTypes = []string{
	"text/javascript",
	"application/javascript",
}

// ErrorMissingCallback is the error for when a callback option is expected but missing.
var ErrorMissingCallback = errors.New("A callback is required for JSONP")

// ErrorUnmarshalNotSupported is the error for when Unmarshal is called but not supported.
var ErrorUnmarshalNotSupported = errors.New("Unmarshalling an object is not supported for JSONP")

// JsonPCodec converts objects to JSONP.
type JsonPCodec struct {
	config Config
	lock   *sync.RWMutex
}

// New returns a new JSONP serializer
func New(cfg ...Config) *JsonPCodec {
	c := DefaultConfig().Merge(cfg)
	return &JsonPCodec{config: c}
}

// Marshal converts an object to JSONP.
func (c *JsonPCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {

	if len(options) == 0 {
		return nil, ErrorMissingCallback
	}

	json, err := jsonEncoding.Marshal(object)

	if err != nil {
		return nil, err
	}

	// #codec-context-options
	// the assumption is options[0] is the callback parameter,
	// and options[1] is the client-context (NB: not *Context) string.

	var callbackFunctionName string
	var callbackString string
	var clientContextString string
	var ok bool

	if callbackFunctionName, ok = options[constants.OptionKeyClientCallback].(string); !ok {
		panic("stretchrcom/codecs: JSONP requires the options to contain the constants.OptionKeyClientCallback value.")
	}

	clientContextString, hasClientContext := options[constants.OptionKeyClientContext].(string)

	if !hasClientContext {
		callbackString = stewstrings.MergeStrings(callbackFunctionName, "(", string(json), ");")
	} else {
		callbackString = stewstrings.MergeStrings(options[constants.OptionKeyClientCallback].(string), "(", string(json), `,"`, clientContextString, `"`, ");")
	}

	return []byte(callbackString), nil
}

// Unmarshal is not supported for JSONP. Returns an error.
func (c *JsonPCodec) Unmarshal(data []byte, obj interface{}) error {
	return ErrorUnmarshalNotSupported
}

// ContentType returns the content type for this codec.
func (c *JsonPCodec) ContentType() string {
	return ContentTypeJSONP
}

// FileExtension returns the file extension for this codec.
func (c *JsonPCodec) FileExtension() string {
	return FileExtensionJSONP
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *JsonPCodec) CanMarshalWithCallback() bool {
	return true
}

func (c *JsonPCodec) ContentTypeSupported(contentType string) bool {
	for _, supportedType := range validJsonpContentTypes {
		if supportedType == contentType {
			return true
		}
	}
	return contentType == c.ContentType()
}
