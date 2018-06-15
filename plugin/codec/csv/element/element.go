package element

import (
	"github.com/sniperkit/textql/pkg/inputs"
	"github.com/sniperkit/textql/pkg/outputs"
)

var (

	// inputOpts
	inputOpts *inputs.CSVInputOptions

	// displayOpts
	displayOpts *outputs.CSVOutputOptions

	// storageSqlite (not implemented as we will use either pivot or gorm as abstraction data layer)
	// storageSqlite *storage.SQLite3Storage
)
