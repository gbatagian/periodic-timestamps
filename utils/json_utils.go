package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJsonString(payload interface{}) string {
	str, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func ToJsonBytesStream(payload interface{}) io.Reader {
	bytes_buffer := new(bytes.Buffer)
	json.NewEncoder(bytes_buffer).Encode(payload)
	return bytes_buffer
}
