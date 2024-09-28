package core

import (
	"encoding/base64"
	"os"
)

func GetSecret() []byte {
	return []byte(base64.URLEncoding.EncodeToString([]byte(os.Getenv("sikretiki"))))
}
