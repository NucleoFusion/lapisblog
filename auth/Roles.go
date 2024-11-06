package auth

import "errors"

type Roles string

const (
	Admin Roles = "Administrator"
	User  Roles = "User"
)

func GetRole(s string) (Roles, error) {
	var role Roles
	var err error = nil

	switch s {
	case "Admin":
		role = Admin
	case "User":
		role = User
	default:
		err = errors.New("invalid role")
	}

	return role, err
}

func (r *Roles) GetString() (string, error) {
	switch *r {
	case Admin:
		return "Admin", nil
	case User:
		return "User", nil
	default:
		return "Error", errors.New("invalid role")
	}
}
