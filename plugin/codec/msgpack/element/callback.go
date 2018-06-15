package element

// Callback is a type alias for OnJSON callback functions
type Callback func(*Element)

// CallbackContainer
type CallbackContainer struct {

	// Query specifies
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies
	Function Callback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	//-- End
}
