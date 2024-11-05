package statics

import enums "lapisblog/statics/Enums"

type Education struct {
	DegreeType           enums.DegreeType  `json:"degreeType"`
	DegreeSpecialization string            `json:"degreeSpecialization"`
	GradeSystem          enums.GradeSystem `json:"gradeSystem"`
	Grade                float32           `json:"grade"`
}
