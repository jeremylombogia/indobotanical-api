package config

import (
	"os"
)

const APPKEY = "5ES2Yt8LNlOmm73nivtlZDAoovnQIRnO"

func GetPort() string {
	return ":" + os.Getenv("PORT")
}
