package filters

import (
	"regexp"

	"github.com/syreclabs/liquid/core"
)

var stripNewLines = &ReplacePattern{regexp.MustCompile("(\n|\r)"), ""}

func StripNewLinesFactory(parameters []core.Value) core.Filter {
	return stripNewLines.Replace
}
