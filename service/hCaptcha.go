package service

import (
	"io"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/jz0ojiang/go-maimai-player-counter/conf"
)

// Generated by https://quicktype.io

type HCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	Credit      bool     `json:"credit"`
	ErrorCodes  []string `json:"error-codes"`
}

func VerifyhCaptcha(hCaptchaResponse string) bool {
	payload := strings.NewReader("response=" + hCaptchaResponse + "&secret=" + conf.GetConfig().GetHCaptchaSecret())
	response, err := http.Post("https://api.hcaptcha.com/siteverify", "application/x-www-form-urlencoded", payload)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}
	var hCaptchaResponseStruct HCaptchaResponse
	err = sonic.Unmarshal(body, &hCaptchaResponseStruct)
	if err != nil {
		return false
	}
	return hCaptchaResponseStruct.Success
}
