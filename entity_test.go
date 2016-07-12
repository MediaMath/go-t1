package t1

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestExecuteRequestError(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:234567", nil)
	c := &http.Client{Timeout: 1 * time.Second}
	if err != nil {
		t.Fatal(err)
	}
	meta, err := execute(req, c, nil)
	if err == nil {
		t.Error("Expected err, got none")
	}
	if want := (Meta{}); !reflect.DeepEqual(meta, want) {
		t.Errorf("request error meta: got %v, want %v", meta, want)
	}
}

// TODO more tests :/
