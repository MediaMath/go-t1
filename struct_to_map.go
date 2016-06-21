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
		tag := type_.Field(i).Tag.Get("url")
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}

		f, v := iVal.Field(i), ""
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(tag, v)
	}

	return values
}
