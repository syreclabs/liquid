package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/syreclabs/liquid/core"
)

func TestReplaceFirstValueInAString(t *testing.T) {
	spec := gspec.New(t)
	filter := ReplaceFirstFactory([]core.Value{stringValue("foo"), stringValue("bar")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barbarforfoo")
}
