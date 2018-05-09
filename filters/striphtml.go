package filters

import (
	"regexp"

	"github.com/syreclabs/liquid/core"
)

var stripHtml = &ReplacePattern{regexp.MustCompile("(?i)<script.*?</script>|<!--.*?-->|<style.*?</style>|<.*?>"), ""}

func StripHtmlFactory(parameters []core.Value) core.Filter {
	return stripHtml.Replace
}
