package utils

import (
	"bytes"
	"encoding/json"
)

func ToJSONReader(v interface{}) *bytes.Reader {
	b, _ := json.Marshal(v)
	return bytes.NewReader(b)
}
