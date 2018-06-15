package element

//go:generate go-enum -f=element_tab.go --marshal --lower
// CODEC x ENUM(JSON,YAML,XML,CSV,TSV,XLSX,ASCII,MARKDOWN,MYSQL,POSTGRES)

// CODEC defines
type CODEC string

const (

	// JSON is the key for `JSON` encoding
	JSON CODEC = "json"

	// YAML is the key for `YAML` encoding
	YAML CODEC = "yaml"

	// XML is the key for `XML` encoding
	XML CODEC = "xml"

	// CSV is the key for `CSV` encoding (columns sperated by comma or semi-colon).
	CSV CODEC = "csv"

	// TSV is the key for `TSV` encoding (tab separated columns).
	TSV CODEC = "tsv"

	// XLSX is the key for `XLSX` encoding. (Microsoft Excel)
	XLSX CODEC = "xlsx"

	// ASCII is the key for `ASCII` encoding
	ASCII CODEC = "ascii"

	// MD is the key for `MARKDOWN` encoding
	MD CODEC = "md"

	// MD is the key for `MARKDOWN` encoding
	MYSQL CODEC = "mysql"

	// POSTGRES is the key for `POSTGRES` encoding
	POSTGRES CODEC = "postgres"

	//-- End
)
