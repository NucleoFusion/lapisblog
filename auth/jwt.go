package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type JWT struct {
	Header    Header
	Payload   Payload
	Signature Signature
}

// Encodes the JWT with Base64 and sha256 encoding
func (jwt *JWT) EncodeJWT() (string, error) {
	payloadData, err := json.Marshal(jwt.Payload)
	if err != nil {
		return "", err
	}

	headerData, err := json.Marshal(jwt.Header)
	if err != nil {
		return "", err
	}

	jws := base64.RawStdEncoding.EncodeToString(headerData) + "." + base64.RawStdEncoding.EncodeToString(payloadData)

	sign, err := ConvertToSignature(jws)
	if err != nil {
		return "", err
	}

	jws += "." + sign

	return jws, nil
}

func DecodeJWT(jws string) (*JWT, error) {
	JWT := JWT{}

	splitJWS := strings.Split(jws, ".")

	err := godotenv.Load(".env")
	if err != nil {
		return &JWT, errors.New(".env file error")
	}

	hmacWriter := hmac.New(sha256.New, []byte(os.Getenv("HMAC_KEY")))
	hmacWriter.Write([]byte(splitJWS[0] + "." + splitJWS[1]))

	hmacNew := hmacWriter.Sum(nil)
	b64Decoded, err := base64.RawStdEncoding.DecodeString(splitJWS[2])
	if err != nil {
		return &JWT, err
	}

	if !hmac.Equal(hmacNew, b64Decoded) {
		return &JWT, errors.New("invalid token, tampering found")
	}

	decodedPayloadData, err := base64.RawStdEncoding.DecodeString(splitJWS[1])
	if err != nil {
		return &JWT, err
	}

	decodedHeaderData, err := base64.RawStdEncoding.DecodeString(splitJWS[0])
	if err != nil {
		return &JWT, err
	}

	err1 := json.Unmarshal(decodedPayloadData, &JWT.Payload)
	err2 := json.Unmarshal(decodedHeaderData, &JWT.Header)
	if err1 != nil {
		return &JWT, err1
	}
	if err2 != nil {
		return &JWT, err1
	}

	JWT.Signature = Signature{Key: splitJWS[2]}

	validityErr := ValidTime(&JWT)
	if validityErr != nil {
		return &JWT, validityErr
	}

	return &JWT, nil
}

func ValidTime(jwt *JWT) error {
	createdAt, _ := jwt.Payload.CreatedAt, jwt.Payload.ValidFor

	fmt.Println(createdAt+validity, time.Now().Unix())

	if createdAt+validity <= time.Now().Unix() {
		return errors.New("expired token")
	}

	return nil
}

// Converts Payload and Header to a JWS Signature with sha256 encoding
func ConvertToSignature(s string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return s, errors.New(".env file error")
	}

	hmacWriter := hmac.New(sha256.New, []byte(os.Getenv("HMAC_KEY")))
	hmacWriter.Write([]byte(s))

	b64Encoded := base64.RawStdEncoding.EncodeToString(hmacWriter.Sum(nil))

	return b64Encoded, nil
}
