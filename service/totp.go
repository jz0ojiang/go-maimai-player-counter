package service

import (
	"github.com/jz0ojiang/go-maimai-player-counter/conf"
	"github.com/pquerna/otp/totp"
)

func ValidateTotp(passcode string) bool {
	return totp.Validate(passcode, conf.GetConfig().GetTotpSecret())
}
