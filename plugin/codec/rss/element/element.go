package element

import (
	models "github.com/sniperkit/codecs/plugin/http/models"
	rssquery "github.com/sniperkit/gofeed/pkg"
)

// RSSElement is the representation of a RSS tag.
type Element struct {

	////
	////// exported /////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string

	// Attrs specifies the attributes to extract
	Attrs []string

	// Request is the request object of the element's HTML document
	Request *models.Request

	// Response is the Response object of the element's HTML document
	Response *models.Response

	// Extractor points to...
	// Extractor *Extractor

	// DOM is the DOM object of the page. DOM is relative
	// to the current RSSElement and is either a html.Node or rssquery.Node
	// based on how the RSSElement was created.
	DOM interface{}

	////
	////// not exported /////////////////////////////////////////////
	////

	// attributes
	attributes interface{}

	//--- End
}

// Callback is a type alias for OnRSS callback functions
type Callback func(*Element)

// CallbackContainer
type CallbackContainer struct {

	// Query specifies
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies
	Function Callback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	//-- End
}

// NewElementFromRSSFeed creates a Element from a rssquery.Feed.
func NewElementFromRSSFeed(resp *models.Response, fp *rssquery.Feed) *Element {

	return &Element{
		Request:  resp.Request,
		Response: resp,
	}
}
