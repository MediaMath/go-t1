package t1

import (
	"net/url"
	"reflect"
	"testing"
)

func TestTaglessElementsSkipped(t *testing.T) {
	type testType struct {
		Int  int
		Int2 int `json:""`
	}
	inp, exp := &testType{Int: 1, Int2: 2}, url.Values{}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap tagless element: want %v, got %v", exp, got)
	}
}

func TestExtendedJSONTagIgnored(t *testing.T) {
	type testType struct {
		Int int `json:"int,omitempty"`
	}
	inp, exp := &testType{Int: 1}, url.Values{"int": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single int: want %v, got %v", exp, got)
	}
}

func TestStructToMapInt(t *testing.T) {
	type testType struct {
		Int int `json:"int"`
	}
	inp, exp := &testType{Int: 1}, url.Values{"int": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single int: want %v, got %v", exp, got)
	}
}

func TestStructToMapInt8(t *testing.T) {
	type testType struct {
		Int8 int8 `json:"int8"`
	}
	inp, exp := &testType{Int8: 1}, url.Values{"int8": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single int8: want %v, got %v", exp, got)
	}
}

func TestStructToMapInt16(t *testing.T) {
	type testType struct {
		Int16 int16 `json:"int16"`
	}
	inp, exp := &testType{Int16: 1}, url.Values{"int16": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single int16: want %v, got %v", exp, got)
	}
}

func TestStructToMapInt32(t *testing.T) {
	type testType struct {
		Int32 int32 `json:"int32"`
	}
	inp, exp := &testType{Int32: 1}, url.Values{"int32": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single int32: want %v, got %v", exp, got)
	}
}

func TestStructToMapInt64(t *testing.T) {
	type testType struct {
		Int64 int64 `json:"int64"`
	}
	inp, exp := &testType{Int64: 1}, url.Values{"int64": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single int64: want %v, got %v", exp, got)
	}
}

func TestStructToMapUint(t *testing.T) {
	type testType struct {
		Uint uint `json:"uint"`
	}
	inp, exp := &testType{Uint: 1}, url.Values{"uint": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single uint: want %v, got %v", exp, got)
	}
}

func TestStructToMapUint8(t *testing.T) {
	type testType struct {
		Uint8 uint8 `json:"uint8"`
	}
	inp, exp := &testType{Uint8: 1}, url.Values{"uint8": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single uint8: want %v, got %v", exp, got)
	}
}

func TestStructToMapUint16(t *testing.T) {
	type testType struct {
		Uint16 uint16 `json:"uint16"`
	}
	inp, exp := &testType{Uint16: 1}, url.Values{"uint16": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single uint16: want %v, got %v", exp, got)
	}
}

func TestStructToMapUint32(t *testing.T) {
	type testType struct {
		Uint32 uint32 `json:"uint32"`
	}
	inp, exp := &testType{Uint32: 1}, url.Values{"uint32": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single uint32: want %v, got %v", exp, got)
	}
}

func TestStructToMapUint64(t *testing.T) {
	type testType struct {
		Uint64 uint64 `json:"uint64"`
	}
	inp, exp := &testType{Uint64: 1}, url.Values{"uint64": []string{"1"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single uint64: want %v, got %v", exp, got)
	}
}

func TestStructToMapFloat32(t *testing.T) {
	type testType struct {
		Float32 float32 `json:"float32"`
	}
	inp, exp := &testType{Float32: 1.0}, url.Values{"float32": []string{"1.0000"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single float32: want %v, got %v", exp, got)
	}
}

func TestStructToMapFloat64(t *testing.T) {
	type testType struct {
		Float64 float64 `json:"float64"`
	}
	inp, exp := &testType{Float64: 1.0}, url.Values{"float64": []string{"1.0000"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single float64: want %v, got %v", exp, got)
	}
}

func TestStructToMapByteSlice(t *testing.T) {
	type testType struct {
		ByteSlice []byte `json:"byte_slice"`
	}
	inp, exp := &testType{ByteSlice: []byte("hi")}, url.Values{"byte_slice": []string{"hi"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single []byte: want %v, got %v", exp, got)
	}
}

func TestStructToMapString(t *testing.T) {
	type testType struct {
		String string `json:"str"`
	}
	inp, exp := &testType{String: "hi"}, url.Values{"str": []string{"hi"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single string: want %v, got %v", exp, got)
	}
}

func TestStructToMapStringSlice(t *testing.T) {
	type testType struct {
		StringSlice []string `json:"string_slice"`
	}
	inp, exp := &testType{StringSlice: []string{"hi", "there"}}, url.Values{"string_slice": []string{"hi", "there"}}
	if got := structToMap(inp); !reflect.DeepEqual(exp, got) {
		t.Errorf("structToMap single []string: want %v, got %v", exp, got)
	}
}
