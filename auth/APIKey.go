package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base32"
	"lapisblog/statics"
	"os"

	"github.com/joho/godotenv"
)

func GetKey(profile *statics.Profile) string {
	godotenv.Load(".env")

	hasher := hmac.New(sha256.New, []byte(os.Getenv("API_KEY")))
	hasher.Write([]byte(profile.Email))
	encoded := base32.StdEncoding.EncodeToString(hasher.Sum(nil))

	return encoded
}
