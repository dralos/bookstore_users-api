package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	has := md5.New()
	defer has.Reset()
	has.Write([]byte(input))
	return hex.EncodeToString(has.Sum(nil))
}
