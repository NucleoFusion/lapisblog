package enums

import "errors"

type Tags string

const (
	Tech  Tags = "Technology"
	Prog  Tags = "Programming"
	Game  Tags = "Gaming"
	Sport Tags = "Sports"
	Music Tags = "Music"
)

func GetTag(str string) (Tags, error) {
	switch str {
	case "Technology":
		return Tech, nil
	case "Programming":
		return Prog, nil
	case "Gaming":
		return Game, nil
	case "Sports":
		return Sport, nil
	case "Music":
		return Music, nil
	}

	return Tags(""), errors.New("tag not found")
}
