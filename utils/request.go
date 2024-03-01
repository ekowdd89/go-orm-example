package utils

import (
	"encoding/json"
	"net/http"
)

func HandlerRequest[T any](w http.ResponseWriter, t *http.Request) (T, error) {
	var req T
	if err := json.NewDecoder(t.Body).Decode(&req); err != nil {
		return req, err
	}
	return req, nil
}
