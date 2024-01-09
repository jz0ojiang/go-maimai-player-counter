package conf

import (
	"os"

	"github.com/bytedance/sonic"
)

type ConfigStruct struct {
	Database        DatabaseStruct `json:"database"`
	HCaptchaSecret  string         `json:"hcaptcha_secret"`
	TurnstileSecret string         `json:"turnstile_secret"`
	Captcha         string         `json:"captcha"`
	TotpSecret      string         `json:"totp_secret"`
	Host            string         `json:"host"`
	Debug           bool           `json:"debug"`
}

type DatabaseStruct struct {
	Sqlite  string `json:"sqlite"`
	Leveldb string `json:"leveldb"`
}

var config ConfigStruct = ConfigStruct{
	Database: DatabaseStruct{
		Sqlite:  "data/pcdata.sqlite3",
		Leveldb: "data/counter",
	},
	Captcha:         "hCaptcha",
	HCaptchaSecret:  "",
	TurnstileSecret: "",
	TotpSecret:      "",
	Host:            ":8080",
	Debug:           false,
}

func init() {
	conf, err := os.ReadFile("config.json")
	if err != nil {
		return
	}
	err = sonic.Unmarshal(conf, &config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() ConfigStruct {
	return config
}

func (ConfigStruct) GetDatabase(t string) string {
	switch t {
	case "sqlite":
		return config.Database.Sqlite
	case "leveldb":
		return config.Database.Leveldb
	default:
		return ""
	}
}

func (ConfigStruct) EditDatabase(t string, v string) {
	switch t {
	case "sqlite":
		config.Database.Sqlite = v
	case "leveldb":
		config.Database.Leveldb = v
	}
}

func (ConfigStruct) GetHCaptchaSecret() string {
	return config.HCaptchaSecret
}

func (ConfigStruct) GetTurnstileSecret() string {
	return config.TurnstileSecret
}

func (ConfigStruct) GetTotpSecret() string {
	return config.TotpSecret
}

func (ConfigStruct) GetHost() string {
	return config.Host
}

func (ConfigStruct) GetDebug() bool {
	return config.Debug
}

func (ConfigStruct) GetCaptcha() string {
	return config.Captcha
}
