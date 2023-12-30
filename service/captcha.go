package service

import "github.com/jz0ojiang/go-maimai-player-counter/conf"

func VerifyCaptcha(captchaResponse string) bool {
	switch conf.GetConfig().Captcha {
	case "hCaptcha":
		return VerifyhCaptcha(captchaResponse)
	case "Turnstile":
		return VerifyTurnstile(captchaResponse)
	default:
		return false
	}
}
