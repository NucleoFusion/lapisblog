package statics

import (
	enums "lapisblog/statics/Enums"
	"time"
)

type Profile struct {
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	BirthDate    time.Time       `json:"birthDate"`
	Education    []Education     `json:"education"`
	LinkedIn     Link            `json:"linkedin"`
	OtherLinks   []Link          `json:"otherLinks"`
	Description  string          `json:"description"`
	Following    []UserReference `json:"following"`
	Followers    []UserReference `json:"followers"`
	TagsFollowed []enums.Tags    `json:"tags"`
}
