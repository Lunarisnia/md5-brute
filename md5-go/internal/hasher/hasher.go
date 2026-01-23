package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(raw string) string {
	hasherBytes := md5.Sum([]byte(raw))
	hashString := hex.EncodeToString(hasherBytes[:])
	return hashString
}
