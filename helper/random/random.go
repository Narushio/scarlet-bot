package random

import (
	"crypto/sha256"
	"time"
)

func Sha256() string {
	hash := sha256.New()
	hash.Write([]byte(time.Now().String()))
	return string(hash.Sum(nil))
}
