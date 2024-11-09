package statics

import (
	enums "lapisblog/statics/Enums"
)

type Profile struct {
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	BirthDate    int64           `json:"birthDate"`
	Education    []Education     `json:"education"`
	LinkedIn     Link            `json:"linkedin"`
	OtherLinks   []Link          `json:"otherLinks"`
	Description  string          `json:"description"`
	Following    []UserReference `json:"following"`
	Followers    []UserReference `json:"followers"`
	TagsFollowed []enums.Tags    `json:"tags"`
}
