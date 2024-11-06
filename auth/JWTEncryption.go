package auth

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"lapisblog/auth/jwt"
)

var key *rsa.PrivateKey

func init() {
	key, _ = jwt.GetKey()
}

func EncryptJWT(JWT *jwt.JWT) ([]byte, error) {

	payloadData, err := json.Marshal(&JWT.Payload)
	if err != nil {
		return []byte("Error"), err
	}

	headerData, err := json.Marshal(&JWT.Header)
	if err != nil {
		return []byte("Error"), err
	}

	data := append(payloadData, []byte(".")...)
	data = append(data, headerData...)

	fmt.Println(string(data))

	encrypted, err := jwt.EncryptMsg(data, key)
	if err != nil {
		return []byte("Error"), err
	}

	fmt.Println(string(encrypted))

	return encrypted, nil
}
