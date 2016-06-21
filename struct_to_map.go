package t1

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func structToMap(data interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(data).Elem()
	type_ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		tag := type_.Field(i).Tag.Get("json")
		if tag == "" {
			continue
		}
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}

		f := iVal.Field(i)
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

	return values
}
