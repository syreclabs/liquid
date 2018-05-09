package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/syreclabs/liquid/core"
)

func TestRemovesValuesFromAString(t *testing.T) {
	spec := gspec.New(t)
	filter := RemoveFactory([]core.Value{stringValue("foo")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barfor")
}
