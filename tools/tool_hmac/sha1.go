package tool_hmac

import (
	"crypto/hmac"
	"crypto/sha1"
)

func HmacSha1(str string, sk string) []byte {
	h := hmac.New(sha1.New, []byte(sk))
	h.Write([]byte(str))
	return h.Sum(nil)
}
