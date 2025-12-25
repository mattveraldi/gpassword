package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var Symbols []string = []string{
	"!",
	"&",
	"-",
	"!",
	"<",
	"@",
	"*",
	"=",
	"\\",
	">",
	"#",
	"(",
	"{",
	":",
	",",
	"$",
	")",
	"}",
	";",
	".",
	"%",
	"_",
	"[",
	"^",
	"+",
	"]",
	"|",
	"'",
	"?",
	"/",
	"~",
	"{",
	"\"\"",
}

const (
	CHARSET string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-={}[]|\\:;\"'<>,.?/`~"
	length  int    = 24
)

func GenerateSecurePassword(symbolsToRemove []string) (string, error) {
	charset := CHARSET
	for _, symbol := range symbolsToRemove {
		charset = strings.ReplaceAll(charset, symbol, "")
	}

	fmt.Print(charset)

	password := ""

	for range length {
		randomIndex, err := GenerateRandomNumber(len(charset))
		if err != nil {
			return "", err
		}
		password += string(charset[randomIndex])
	}

	return password, nil
}

func GenerateRandomNumber(length int) (int64, error) {
	interval := big.NewInt(int64(length))
	randomIndex, err := rand.Int(rand.Reader, interval)
	if err != nil {
		return -1, err
	}
	return randomIndex.Int64(), nil
}
