package element

import (
	"strings"

	models "github.com/sniperkit/codecs/plugin/http/models"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// HTMLElement is the representation of a HTML tag.
type HTMLElement struct {

	////
	////// exported /////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string
	// Text is...
	Text string

	// Extractor
	// Extractor *Extractor

	// Request is the request object of the element's HTML document
	Request *models.Request

	// Response is the Response object of the element's HTML document
	Response *models.Response

	// DOM is the goquery parsed DOM object of the page. DOM is relative
	// to the current HTMLElement
	DOM *goquery.Selection

	////
	////// not exported /////////////////////////////////////////////
	////

	// attributes is a list of html attrs
	attributes []html.Attribute
}

// NewHTMLElementFromSelectionNode creates a Element from a goquery.Selection Node.
func NewHTMLElementFromSelectionNode(resp *models.Response, s *goquery.Selection, n *html.Node) *Element {
	return &HTMLElement{
		Name:       n.Data,
		Request:    resp.Request,
		Response:   resp,
		Text:       goquery.NewDocumentFromNode(n).Text(),
		DOM:        s,
		attributes: n.Attr,
	}
}

/*
func (h *Element) Name() string {
	return h.Name
}

func (h *Element) Text() string {
	return h.Text
}
*/

// Attr returns the selected attribute of a Element or empty string
// if no attribute found
func (h *Element) Attr(k string) string {
	for _, a := range h.attributes {
		if a.Key == k {
			return a.Val
		}
	}
	return ""
}

// HTML: Attr, ChildText, ChildAttr, ChildAttrs, ForEach

// ChildText returns the concatenated and stripped text content of the matching elements.
func (h *Element) ChildText(goquerySelector string) string {
	return strings.TrimSpace(h.DOM.Find(goquerySelector).Text())
}

// ChildAttr returns the stripped text content of the first matching element's attribute.
func (h *Element) ChildAttr(goquerySelector, attrName string) string {
	if attr, ok := h.DOM.Find(goquerySelector).Attr(attrName); ok {
		return strings.TrimSpace(attr)
	}
	return ""
}

// ChildAttrs returns the stripped text content of all the matching element's attributes.
func (h *Element) ChildAttrs(goquerySelector, attrName string) []string {
	var res []string
	h.DOM.Find(goquerySelector).Each(func(_ int, s *goquery.Selection) {
		if attr, ok := s.Attr(attrName); ok {
			res = append(res, strings.TrimSpace(attr))
		}
	})
	return res
}

// ForEach iterates over the elements matched by the first argument and calls the callback function on every Element match.
func (h *Element) ForEach(goquerySelector string, callback func(int, *Element)) {
	i := 0
	h.DOM.Find(goquerySelector).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			callback(i, NewHTMLElementFromSelectionNode(h.Response, s, n))
			i++
		}
	})
}
