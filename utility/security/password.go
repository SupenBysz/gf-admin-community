package security

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

// GenerateRandomSalt 生成随机盐值
func GenerateRandomSalt() (string, error) {
	bytes := make([]byte, 16) // 16字节 = 32字符的十六进制字符串
	if _, err := rand.Read(bytes); err != nil {
		return "", gerror.Wrap(err, "生成随机盐值失败")
	}
	return hex.EncodeToString(bytes), nil
}

// CryptoPasswordWithRandomSalt 使用随机盐值和bcrypt加密密码
func CryptoPasswordWithRandomSalt(ctx context.Context, passwordStr string) (hashedPassword, salt string, err error) {
	// 生成随机盐值
	salt, err = GenerateRandomSalt()
	if err != nil {
		return "", "", err
	}
	
	// 将密码和盐值组合
	passwordWithSalt := passwordStr + salt
	
	// 使用bcrypt加密
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", "", gerror.Wrap(err, "密码加密失败")
	}
	
	return string(hashedBytes), salt, nil
}

// VerifyPasswordWithSalt 验证密码（使用盐值）
func VerifyPasswordWithSalt(ctx context.Context, passwordStr, hashedPassword, salt string) error {
	// 将输入密码和盐值组合
	passwordWithSalt := passwordStr + salt
	
	// 使用bcrypt验证
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithSalt))
	if err != nil {
		return gerror.Wrap(err, "密码验证失败")
	}
	
	return nil
}

// CryptoPasswordWithFixedSalt 兼容旧版本的固定盐值加密（用于迁移）
func CryptoPasswordWithFixedSalt(ctx context.Context, passwordStr, salt string) (string, error) {
	// 将密码和盐值组合
	passwordWithSalt := passwordStr + salt
	
	// 使用bcrypt加密
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", gerror.Wrap(err, "密码加密失败")
	}
	
	return string(hashedBytes), nil
}

// IsPasswordNeedUpgrade 检查密码是否需要升级（从固定盐值升级到随机盐值）
func IsPasswordNeedUpgrade(ctx context.Context, user interface{}) bool {
	// 这里可以根据具体的用户结构来判断
	// 例如：如果salt字段为空或者使用的是旧的加密方式，则需要升级
	// 具体实现可以根据业务需求调整
	return false
}

// MigratePassword 迁移密码（从旧加密方式升级到新加密方式）
func MigratePassword(ctx context.Context, oldPassword, newPassword string) (hashedPassword, salt string, err error) {
	// 验证旧密码是否正确（这里需要根据具体的旧加密方式实现）
	// ...
	
	// 使用新的加密方式加密新密码
	return CryptoPasswordWithRandomSalt(ctx, newPassword)
}

// PasswordStrength 密码强度检查
func PasswordStrength(password string) (score int, suggestions []string) {
	score = 0
	suggestions = make([]string, 0)
	
	if len(password) < 8 {
		suggestions = append(suggestions, "密码长度至少8位")
	} else {
		score += 20
	}
	
	if len(password) >= 12 {
		score += 10
	}
	
	// 检查是否包含数字
	hasNumber := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}
	if hasNumber {
		score += 20
	} else {
		suggestions = append(suggestions, "密码应包含数字")
	}
	
	// 检查是否包含小写字母
	hasLower := false
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
			break
		}
	}
	if hasLower {
		score += 20
	} else {
		suggestions = append(suggestions, "密码应包含小写字母")
	}
	
	// 检查是否包含大写字母
	hasUpper := false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
			break
		}
	}
	if hasUpper {
		score += 20
	} else {
		suggestions = append(suggestions, "密码应包含大写字母")
	}
	
	// 检查是否包含特殊字符
	hasSpecial := false
	specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?"
	for _, char := range password {
		for _, special := range specialChars {
			if char == special {
				hasSpecial = true
				break
			}
		}
		if hasSpecial {
			break
		}
	}
	if hasSpecial {
		score += 10
	} else {
		suggestions = append(suggestions, "密码应包含特殊字符")
	}
	
	return score, suggestions
}