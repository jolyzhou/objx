package objx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	objxStringPrefix = "•"
	objxStringSuffix = "•"
)

// Obj gets the underlying object contained within this
// O.
func (o *O) Obj() interface{} {
	return o.obj
}

// String gets a string representation of the object
// contained.
func (o *O) String() string {

	// if the object has a string method, just call it
	if s, ok := o.obj.(interface {
		String() string
	}); ok {
		return objxStringPrefix + s.String() + objxStringSuffix
	}

	// otherwise, let fmt do the work
	return fmt.Sprintf(objxStringPrefix+"%v"+objxStringSuffix, o.obj)

}

// arrayAccesRegexString is the regex used to extract the array number
// from the access path
const arrayAccesRegexString = `^(.+)\[([0-9]+)\]$`

// arrayAccesRegex is the compiled arrayAccesRegexString
var arrayAccesRegex = regexp.MustCompile(arrayAccesRegexString)

// Get gets a value from this object.
func (o *O) Get(selector interface{}) *O {

	switch selector.(type) {
	case string: // "address.postcode.inner"

		selStr := selector.(string)
		segs := strings.Split(selStr, ".")

		current := o.Obj()

		for _, field := range segs {

			arrayMatches := arrayAccesRegex.FindStringSubmatch(field)

			if len(arrayMatches) > 0 {
				// An array notation has been detected.

				// Get the key into the map
				mName := arrayMatches[1]

				// Get the index into the array at the key
				index, err := strconv.Atoi(arrayMatches[2])
				if err != nil {
					// This should never happen. If it does, something has gone
					// seriously wrong. Panic.
					panic("objx: GEt - array index is not an integer. This should never happen.")
				}

				if m, ok := current.(map[string]interface{}); ok {
					if a, ok := m[mName].([]interface{}); ok {
						current = a[index]
					}
				}
			} else {
				if m, ok := current.(map[string]interface{}); ok {
					current = m[field]
				}
			}

		}

		return New(current)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:

		selInt := uint64FromInterface(selector)
		if a, ok := o.Obj().([]interface{}); ok {
			return New(a[selInt])
		}

	}

	panic("objx: Get with invalid selector.")

}

// uint64FromInterface converts an interface object to the largest
// representation of an unsigned integer using a type switch and
// assertions
func uint64FromInterface(selector interface{}) uint64 {
	var value uint64
	switch selector.(type) {
	case int:
		value = uint64(selector.(int))
	case int8:
		value = uint64(selector.(int8))
	case int16:
		value = uint64(selector.(int16))
	case int32:
		value = uint64(selector.(int32))
	case int64:
		value = uint64(selector.(int64))
	case uint:
		value = uint64(selector.(uint))
	case uint8:
		value = uint64(selector.(uint8))
	case uint16:
		value = uint64(selector.(uint16))
	case uint32:
		value = uint64(selector.(uint32))
	case uint64:
		value = selector.(uint64)
	default:
		panic("objx: array access argument is not an integer (this should never happen)")
	}

	return value
}
