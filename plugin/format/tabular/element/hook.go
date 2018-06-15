package element

// Hooks specifies
type Hooks struct {

	// Enabled
	Enabled bool `default:"true" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`

	// Registry
	Registry map[string]*Hook `json:"registry" yaml:"registry" toml:"registry" xml:"-" ini:"registry" csv:"registry"`
}
