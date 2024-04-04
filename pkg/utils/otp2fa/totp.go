package otp2fa

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateTotp(email string) (*otp.Key, error) {
	//totp:Time-Based One-Time Password
	return totp.Generate(totp.GenerateOpts{
		Issuer:      "bibi-demo",
		AccountName: email,
		Period:      300,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
}

func CheckTotp(Totp string, dbTotp string) bool {
	return totp.Validate(Totp, dbTotp)
}
