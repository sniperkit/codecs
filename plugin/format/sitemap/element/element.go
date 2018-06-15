package element

import (
	models "github.com/sniperkit/codecs/plugin/http/models"
	tabular "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

// Callback is a type alias for On callback functions
type Callback func(*Element)

// tabCallbackContainer specifies the struct
type CallbackContainer struct {

	// Query specifies...
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Hook specifies...
	Hook *Hook `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies...
	Function Callback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// Element is the representation of a  tag.
type Element struct {

	////
	////// exported //////////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Event is the name of the pre-processing task to trigger
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Dataset represents...
	Dataset *tabular.Dataset `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Hooks represents...
	Hook *Hook `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Extractor
	// Extractor *Extractor `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Text
	Text string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Request is the request object of the element's HTML document
	Request *models.Request `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Response is the Response object of the element's HTML document
	Response *models.Response `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	////
	////// not exported /////////////////////////////////////////////
	////

	// err stores the loading error
	errs []error `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// NewElementFromNode instanciates a new Element extracted with a tabular data query
func NewElementFromNode(resp *models.Response, query string, ds *tabular.Dataset) *Element {
	t := &Element{
		Dataset:  ds,
		Name:     query,
		Request:  resp.Request,
		Response: resp,
	}
	return t
}

// NewElementFromSelect instanciates a new Element extracted with a global processing hook
func NewElementFromSelect(resp *models.Response, hook *Hook, ds *tabular.Dataset) *Element {
	t := &Element{
		Dataset:  ds,
		Hook:     hook,
		Request:  resp.Request,
		Response: resp,
	}
	return t
}
