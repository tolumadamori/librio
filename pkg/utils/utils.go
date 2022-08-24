package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// Parses the request body and returns byte slice.
func ParseBody(r *http.Request, x interface{}) []byte {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return []byte(err.Error())
		}

	}

	return nil

}
