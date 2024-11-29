package utils

import "net/url"

func FormatURLParameters(inputs map[string]string) string {
	values := url.Values{}
	for k, v := range inputs {
		if v != "" {
			values.Add(k, v)
		}
	}
	return "?" + values.Encode()
}
