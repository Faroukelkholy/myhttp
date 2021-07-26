package myhttp

import (
	"crypto/md5"
	"fmt"
)

func HashMD5(data []byte) string {
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}
