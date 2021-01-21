package service

import (
	"io/ioutil"
)

type Config struct {
	ClientID string
	SecretKey string
	Scopes []string
	UserAgent string
	RedirectUrl string
}

var _config *Config

func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func WriteFile(path string, content []byte) (bool, error) {
	err := ioutil.WriteFile(path, content, 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}

//func InitConfig() {
//	scopes := strings.Split(os.Getenv("EVE_ESI_SCOPES"), ",")
//
//	_config = &Config{
//		ClientID:  os.Getenv("EVE_ESI_CLIENT_ID"),
//		SecretKey: os.Getenv("EVE_ESI_SECRET_KEY"),
//		Scopes:    scopes,
//		RedirectUrl: os.Getenv("EVE_ESI_REDIRECT_URL"),
//		UserAgent: os.Getenv("EVE_ESI_USERAGENT"),
//	}
//}

func GetConfig() *Config {
	return _config
}