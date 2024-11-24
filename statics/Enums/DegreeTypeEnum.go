package enums

import "errors"

type DegreeType string

const (
	Bachelors DegreeType = "Bachelors"
	Masters   DegreeType = "Masters"
	Phd       DegreeType = "Phd"
	School    DegreeType = "School"
)

func GetDegree(s string) (DegreeType, error) {
	switch s {
	case "Bachelors":
		return Bachelors, nil
	case "Masters":
		return Masters, nil
	case "Phd":
		return Phd, nil
	case "School":
		return School, nil
	}

	return DegreeType(""), errors.New("invalid Degree, only Bachelors, Masters, Phd and School accepted")
}
