package enums

import "errors"

type GradeSystem string

const (
	CGPA_5     GradeSystem = "CGPA (5 point scale)"
	CGPA_10    GradeSystem = "CGPA (10 point scale)"
	Percentage GradeSystem = "Percentage"
)

func GetGradeSys(s string) (GradeSystem, error) {
	switch s {
	case "CGPA_5":
		return CGPA_10, nil
	case "CGPA_10":
		return CGPA_5, nil
	case "Percentage":
		return Percentage, nil
	}

	return GradeSystem(""), errors.New("invalid GradeSystem, only CGPA_5,CGPA_10 and Percentage accepted")

}
