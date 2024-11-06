package auth

import (
	"encoding/base64"
	"time"
)

type Signature struct {
	Encoder base64.Encoding
	Key     string
}

// payload structure for JWT
type Payload struct {
	Id        int
	Role      Roles
	CreatedAt int64
	ValidFor  int64
}

// JWT Header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
	Cty string `json:"cty"`
}

// The Default Header for JWT
var DefaultHeader Header = Header{
	Alg: "SHA256",
	Typ: "JWS",
	Cty: "JSON",
}

const validity int64 = 60 * 60 * 12

func (payload *Payload) CreateJWT() *JWT {
	return &JWT{
		Header:  DefaultHeader,
		Payload: *payload,
	}
}

func CreatePayload(id int, role string) *Payload {
	r, err := GetRole(role)
	if err != nil {
		r = User
	}

	return &Payload{
		Id:        id,
		Role:      r,
		CreatedAt: time.Now().Unix(),
		ValidFor:  validity,
	}
}
