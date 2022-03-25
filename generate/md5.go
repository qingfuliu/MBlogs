package generate

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

func Md5(str string) (string, error) {
	w := md5.New()
	_, err := io.WriteString(w, str)
	if err != nil {
		return "", errors.New("md5 failed")
	}
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str, nil
}
