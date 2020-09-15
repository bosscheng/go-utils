package utils

import "strings"

var replacer = strings.NewReplacer("-", "")

func NewID() string {
	return replacer.Replace(NewV4().String())
}
