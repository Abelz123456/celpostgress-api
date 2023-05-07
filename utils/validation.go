package utils

import (
	"crypto/sha1"
)

func EncodeSHA1(messages []string) []byte {
	h := sha1.New()

	if len(messages) > 0 {
		for _, message := range messages {
			h.Write([]byte(message))
		}
	}

	return h.Sum(nil)
}
