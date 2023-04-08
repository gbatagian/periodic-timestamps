package utils

import (
	"io"
	"strings"
	"testing"
)

func TestToJsonString(t *testing.T) {
	// Arrange
	sample_map := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": 1.5999,
		"key4": false,
		"key5": nil,
	}

	// Act
	jsonString := ToJsonString(sample_map)

	// Assert
	if jsonString != `{"key1":"value1","key2":1,"key3":1.5999,"key4":false,"key5":null}` {
		t.Errorf("Invalid conversion to json string: %s", jsonString)
	}
}

func TestToJsonBytesStream(t *testing.T) {
	// Arrange
	sample_map := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": 1.5999,
		"key4": false,
		"key5": nil,
	}

	// Act
	jsonBytesStream := ToJsonBytesStream(sample_map)
	buffer := new(strings.Builder)
	io.Copy(buffer, jsonBytesStream)

	// Assert
	if s := strings.Trim(buffer.String(), "\n"); s != `{"key1":"value1","key2":1,"key3":1.5999,"key4":false,"key5":null}` {
		t.Errorf("Invalid conversion to json bytes stream: %s", s)
	}
}
