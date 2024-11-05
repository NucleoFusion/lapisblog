package jwt

import (
	"crypto/rand"
	"crypto/rsa"
)

func GetKey() (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, 256)
	if err != nil {
		return key, err
	}

	return key, nil
}

func EncryptMsg(data []byte, key *rsa.PrivateKey) ([]byte, error) {
	d, err := rsa.EncryptPKCS1v15(rand.Reader, &key.PublicKey, data)

	return d, err
}

func DecryptMsg(data []byte, key *rsa.PrivateKey) ([]byte, error) {
	d, err := rsa.DecryptPKCS1v15(rand.Reader, key, data)

	return d, err
}
