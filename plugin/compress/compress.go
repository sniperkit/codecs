package compress

import (
	"fmt"

	compressible "github.com/sniperkit/compressible/pkg"
)

func IsCompressible(contentType string) bool {
	return compressible.Is(contentType)
}

func IsCompressibleWithThreshold(contentType string, threshold int) bool {
	var wt compressible.WithThreshold = threshold
	return wt.Compressible(contentType, 1023)
}
