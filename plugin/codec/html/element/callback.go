package element

// Callback is a type alias for OnHTML callback functions
type Callback func(*Element)

// htmlCallbackContainer
type CallbackContainer struct {

	// Selector specifies...
	Selector string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies...
	Function Callback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	//-- End
}
