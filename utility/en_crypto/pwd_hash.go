package en_crypto

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"golang.org/x/crypto/scrypt"
)

// PwdHash 将指定的明文 + 盐进行hash加密
func PwdHash(text string, salt string) (string, error) {
	// 判断盐是否足够8位
	saltLen := len(salt)
	for saltLen < 8 {
		salt += "0"
		saltLen++
	}

	salt = gstr.SubStr(salt, saltLen-8, 8)

	key, err := scrypt.Key([]byte(text), []byte(salt), 1<<15, 8, 1, 32)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}
