package ktalk

import "os"

const (
	ApiURL = "https://%s.ktalk.ru/api"
)

func GetToken() string {
	return os.Getenv("KTALK_TOKEN")
}
