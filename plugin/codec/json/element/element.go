package element

import (
	"strings"

	models "github.com/sniperkit/codecs/plugin/http/models"
	jsonquery "github.com/sniperkit/colly/plugins/data/extract/query/json"
	jsonql_v1 "github.com/sniperkit/jsonql/pkg"
)

// Element is the representation of a JSON tag.
type Element struct {

	////
	////// exported /////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string
	// Text is the content node
	Text string

	// Request is the request object of the element's HTML document
	Request *models.Request

	// Response is the Response object of the element's HTML document
	Response *models.Response

	// Extractor
	// Extractor *Extractor

	// DOM is the DOM object of the page. DOM is relative
	// to the current Element and is either a html.Node or jsonquery.Node
	// based on how the Element was created.
	DOM interface{}

	////
	////// not exported /////////////////////////////////////////////
	////

	// jsonql_v1 --> github.com/elgs/jsonql
	jsonql_v1 *jsonql_v1.JSONQL

	//-- End
}

// NewElementFromJSONNode creates a Element from a jsonquery.Node.
func NewElementFromJSONNode(resp *Response, s *jsonquery.Node) *Element {
	return &Element{
		Name:     s.Data,
		Request:  resp.Request,
		Response: resp,
		Text:     s.InnerText(),
		DOM:      s,
	}
}

// JSON: Find, FindOne, Extract, Header, Headers, Keys, Values, Map,

// Extract
func (h *Element) Extract(pluckerConfig map[string]interface{}) string {
	return ""
}

// Header
func (h *Element) Header(key string) (value string) {
	value = strings.TrimSpace(h.Response.Headers.Get(key))
	return
}

// Headers
func (h *Element) Headers() map[string]string {
	res := make(map[string]string, 0)
	/*
		for key, val := range h.Response.Headers {
			res[key] = val
		}
	*/
	return res
}

// FindOne
func (h *Element) FindOne(xpathQuery string) string {

	n := jsonquery.FindOne(h.DOM.(*jsonquery.Node), xpathQuery)
	if n == nil {
		return ""
	}
	return strings.TrimSpace(n.InnerText())
}

// Find
func (h *Element) Find(xpathQuery, attrName string) []string {
	var res []string
	child := jsonquery.Find(h.DOM.(*jsonquery.Node), xpathQuery)
	for _, node := range child {
		res = append(res, node.InnerText())
	}
	return res
}
