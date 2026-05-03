package service

import (
	"regexp"
)

func IsMorse(text string) bool {
	isMorse, _ := regexp.MatchString(`^[.\-\s/]+$`, string(text))
	return isMorse
}
