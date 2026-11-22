package main

import (
	"strings"
)

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(p []byte) (n int, err error) {
	u.UpperString = strings.ToUpper(string(p))
	return len(p), nil
}
