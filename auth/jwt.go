package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type JWT struct {
	Header    Header
	Payload   Payload
	Signature Signature
}

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

// Converts Payload and Header to a JWS Signature with sha256 encoding
func ConvertToSignature(s string) (string, error) {
	shaEncoded := sha256.Sum256([]byte(s))

	b64Encoded := base64.RawStdEncoding.EncodeToString(shaEncoded[:])

	return b64Encoded, nil
}
