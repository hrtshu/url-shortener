package core

import (
	"github.com/jmcvetta/randutil"
	"net/url"
)

func generateId(size int) (string, error) {
	const lowerAlphabet = "abcdefghijklmnopqrstuvwxyz"
	charset := lowerAlphabet + randutil.Numerals
	id, err := randutil.String(size, charset)
	return id, err
}

func isValidUrl(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
