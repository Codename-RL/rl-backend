package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"math/big"
	"strings"
	"time"
)

var digits = "0123456789"

// GenerateNumericOTP returns a cryptographically secure numeric OTP of the given length.
func GenerateNumericOTP(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	var sb strings.Builder
	sb.Grow(length)
	maximumDigit := big.NewInt(int64(len(digits)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, maximumDigit)
		if err != nil {
			return "", err
		}
		sb.WriteByte(digits[n.Int64()])
	}

	return sb.String(), nil
}

// GenerateOTPWithExpiry generates an OTP and returns it along with the expiry time.
// Pass any custom ttl (e.g. 5*time.Minute).
func GenerateOTPWithExpiry(length int, ttl time.Duration) (string, time.Time, error) {
	code, err := GenerateNumericOTP(length)
	if err != nil {
		return "", time.Time{}, err
	}
	return code, time.Now().Add(ttl), nil
}

// ValidateOTP performs a constant-time comparison of expected vs provided OTP.
func ValidateOTP(expected, provided string) bool {
	if expected == "" || provided == "" {
		return false
	}
	// constant time compare requires equal length
	if len(expected) != len(provided) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(expected), []byte(provided)) == 1
}
