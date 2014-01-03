package filters

import (
	"reflect"
	"github.com/karlseguin/liquid/core"
)

// Creates a first filter
func FirstFactory(parameters []core.Value) Filter {
	return First
}

// get the first element of the passed in array
func First(input interface{}, data map[string]interface{}) interface{} {
	value := reflect.ValueOf(input)
	kind := value.Kind()
	if (kind != reflect.Array && kind != reflect.Slice) || value.Len() == 0 {
		return input
	}
	return value.Index(0).Interface()
}
