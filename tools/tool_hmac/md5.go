package tool_hmac

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

// Md5String .
func Md5String(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Md5Bytes .
func Md5Bytes(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// FileMd5Sum .
func FileMd5Sum(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(data)), nil
}
