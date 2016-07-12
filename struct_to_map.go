package t1

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func structToMapGivenValues(data interface{}, values url.Values) {
	val := reflect.ValueOf(data).Elem()
	valType := val.Type()
	for i := 0; i < val.NumField(); i++ {
		tag := valType.Field(i).Tag.Get("json")
		if tag == "" {
			continue
		}
		var opts tagOptions
		tag, opts = parseTag(tag)

		f := val.Field(i)
		if !f.IsValid() || opts.Contains("omitempty") && isEmptyValue(f) {
			continue
		}

		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			values.Add(tag, strconv.FormatInt(f.Int(), 10))
		case uint, uint8, uint16, uint32, uint64:
			values.Add(tag, strconv.FormatUint(f.Uint(), 10))
		case float32:
			values.Add(tag, strconv.FormatFloat(f.Float(), 'f', 4, 32))
		case float64:
			values.Add(tag, strconv.FormatFloat(f.Float(), 'f', 4, 64))
		case []byte:
			values.Add(tag, string(f.Bytes()))
		case string:
			values.Add(tag, f.String())
		case []string:
			for j := 0; j < f.Len(); j++ {
				values.Add(tag, f.Index(j).String())
			}
		}
	}
}

func structToMap(data interface{}) url.Values {
	values := make(url.Values)
	structToMapGivenValues(data, values)
	return values
}

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
