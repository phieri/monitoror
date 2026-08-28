package humanize

import (
	"fmt"
	"reflect"

	"github.com/dustin/go-humanize"
)

// Interface transform interface to string
func Interface(value interface{}) string {
	if value == nil {
		return ""
	}

	valueOf := reflect.ValueOf(value)
	for valueOf.IsValid() && valueOf.Kind() == reflect.Ptr {
		if valueOf.IsNil() {
			return ""
		}
		valueOf = valueOf.Elem()
	}

	if !valueOf.IsValid() {
		return ""
	}

	value = valueOf.Interface()

	switch value := value.(type) {
	case float32:
		return humanize.Ftoa(float64(value))
	case float64:
		return humanize.Ftoa(value)
	default:
		return fmt.Sprintf("%v", value)
	}
}
