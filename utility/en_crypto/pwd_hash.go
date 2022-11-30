package en_crypto

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

// PwdEncodeHash 将指定的明文 + 盐进行hash加密
func PwdEncodeHash(text []byte, salt []byte) (string, error) {
	key, err := scrypt.Key(text, salt, 1<<15, 8, 1, 32)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}
