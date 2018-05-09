package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/syreclabs/liquid/core"
)

func TestRemovesFirstValueFromAString(t *testing.T) {
	spec := gspec.New(t)
	filter := RemoveFirstFactory([]core.Value{stringValue("foo")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barforfoo")
}
