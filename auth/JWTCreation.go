package auth

import (
	"lapisblog/auth/jwt"
	"time"
)

// Creates a Payload with the provided id, role and validity.
// For Default validity put validity = -1 in function call.
// Returns an error if role is invalid, currently accepted roles
// are "Admin" and "User"
func CreatePayload(id int, role string, validity int64) (*jwt.Payload, error) {
	if validity == -1 {
		validity = 60 * 60 * 6
	}

	r, err := GetRole(role)
	if err != nil {
		return &jwt.Payload{}, err
	}

	s, err := r.GetString()
	if err != nil {
		return &jwt.Payload{}, err
	}

	payload := jwt.Payload{
		UserId:    id,
		Role:      s,
		CreatedAt: time.Now().UnixNano(),
		ValidFor:  validity,
	}

	return &payload, nil
}

// Creates a JWT with the provided payload.
// The JWT has RSA Encrption
func GetJWT(payload *jwt.Payload) *jwt.JWT {
	jwtStruct := jwt.JWT{
		Header: jwt.Header{
			Alg: "RSA",
			Typ: "JWT",
		},
		Payload: *payload,
	}

	return &jwtStruct
}
