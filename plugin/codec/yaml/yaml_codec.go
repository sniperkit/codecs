package yaml

import (
	"sync"

	yamlEncoding "github.com/go-yaml/yaml"
)

const (
	ContentTypeYAML   string = "application/yaml"
	FileExtensionYAML string = ".yaml"
)

var validYamlContentTypes = []string{
	"application/yaml",
	"text/yaml",
}

// YamlCodec converts objects to and from YAML.
type YamlCodec struct {
	config Config
	lock   *sync.RWMutex
}

// New returns a new YAML serializer
func New(cfg ...Config) *YamlCodec {
	c := DefaultConfig().Merge(cfg)
	return &YamlCodec{config: c}
}

// Converts an object to YAML.
func (c *YamlCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return yamlEncoding.Marshal(object)
}

// Unmarshal converts YAML into an object.
func (c *YamlCodec) Unmarshal(data []byte, obj interface{}) error {
	return yamlEncoding.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (c *YamlCodec) ContentType() string {
	return ContentTypeYAML
}

// FileExtension returns the file extension for this codec.
func (c *YamlCodec) FileExtension() string {
	return FileExtensionYAML
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *YamlCodec) CanMarshalWithCallback() bool {
	return false
}

func (c *YamlCodec) ContentTypeSupported(contentType string) bool {
	for _, supportedType := range validYamlContentTypes {
		if supportedType == contentType {
			return true
		}
	}
	return contentType == c.ContentType()
}
