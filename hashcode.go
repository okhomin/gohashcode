package gohashcode

import (
	"math"
	"reflect"
	"slices"
	"strings"
)

var (
	hashcoderType = reflect.TypeFor[Hashcoder]()
)

// Hashcoder is the interface implemented by types that
// can compute their own valid hashcode.
type Hashcoder interface {
	Hashcode() uint64
}

// Hashcode returns a hashcode of v
// Hashcode traverses the value v recursively. If an encountered value implements Hashcoder and is not a nil pointer,
// Hashcode calls [Hashcoder.Hashcode] to compute the hashcode of the value. Otherwise, Hashcode uses the default hashcode algorithm.
func Hashcode(v any) uint64 {
	// if the value is nil, return 0
	if v == nil {
		return 0
	}

	t := reflect.TypeOf(v)

	// if the value implements Hashcoder, call Hashcoder.Hashcode
	if t.Implements(hashcoderType) {
		return v.(Hashcoder).Hashcode()
	}

	// switch the type of the value
	switch t.Kind() {
	case reflect.Bool:
		return boolHashcode(v.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intHashcode(uint64(reflect.ValueOf(v).Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return intHashcode(reflect.ValueOf(v).Uint())
	case reflect.Float32, reflect.Float64:
		return floatHashcode(reflect.ValueOf(v).Float())
	case reflect.String:
		return sliceHashcode([]byte(reflect.ValueOf(v).String()))
	case reflect.Slice, reflect.Array:
		hash := uint64(7)
		for i := 0; i < reflect.ValueOf(v).Len(); i++ {
			hash = 31*hash + Hashcode(reflect.ValueOf(v).Index(i).Interface())
		}
		return hash
	case reflect.Map:
		hash := uint64(7)

		keys := reflect.ValueOf(v).MapKeys()
		// sort the keys to ensure the hashcode is consistent
		slices.SortFunc(keys, func(a, b reflect.Value) int {
			return strings.Compare(a.String(), b.String())
		})
		for _, k := range keys {
			hash = 31*hash + Hashcode(k.Interface())
			hash = 31*hash + Hashcode(reflect.ValueOf(v).MapIndex(k).Interface())
		}
		return hash
	case reflect.Struct:
		return structHashcode(v)
	case reflect.Func:
		// signature of the function
		return sliceHashcode([]byte(t.String()))
	case reflect.Ptr:
		if !reflect.ValueOf(v).IsNil() {
			return Hashcode(reflect.ValueOf(v).Elem().Interface())
		}
	}
	return 0
}

func boolHashcode(v bool) uint64 {
	if v {
		return 1231
	}
	return 1237
}

func intHashcode(v uint64) uint64 {
	return 31 * uint64(v)
}

func floatHashcode(v float64) uint64 {
	return 31 * math.Float64bits(v)
}

func sliceHashcode[T any](v []T) uint64 {
	hash := uint64(7)
	for _, e := range v {
		hash = 31*hash + Hashcode(e)
	}
	return hash
}

func structHashcode(v any) uint64 {
	// 7 + hash of the type name
	hash := uint64(7) + sliceHashcode([]byte(reflect.TypeOf(v).Name()))
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() || f.Tag.Get("hash") == "false" || f.Tag.Get("hash") == "-" {
			continue
		}

		hash = 31*hash + Hashcode(reflect.ValueOf(v).Field(i).Interface())
	}
	return hash
}
