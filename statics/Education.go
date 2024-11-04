package statics

import enums "lapisblog/statics/Enums"

type Education struct {
	Degree      enums.DegreeType  `json:"degree"`
	Name        string            `json:"name"`
	GradeSystem enums.GradeSystem `json:"gradeSystem"`
	Grade       float32           `json:"grade"`
}
