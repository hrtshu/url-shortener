package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func extractIdFromHostname(hostname string, baseHostname string, idSize int) (string, error) {
	if !strings.HasSuffix(hostname, baseHostname) {
		return "", errors.New("invalid url")
	}
	prefix := hostname[:len(hostname)-len(baseHostname)]
	if len(prefix) < 2 {
		return "", errors.New("invalid url")
	}
	if prefix[len(prefix)-1] != '.' {
		return "", errors.New("invalid url")
	}
	prefix = prefix[:len(prefix)-1]
	if strings.Count(prefix, ".") != 0 {
		return "", errors.New("invalid url")
	}
	if len(prefix) != idSize {
		return "", errors.New("invalid url")
	}
	// Note: The character type and length are not checked here.
	return prefix, nil
}

func respondJson(w http.ResponseWriter, resp interface{}) error {
	res, err := json.Marshal(resp)
	if err != nil {
		return err // TODO
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return nil
}

func redirect(w http.ResponseWriter, url string) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", url)
	w.WriteHeader(http.StatusMovedPermanently)
}
