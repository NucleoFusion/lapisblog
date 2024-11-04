package main

import (
	"encoding/json"
	"fmt"
	"lapisblog/statics"
	enums "lapisblog/statics/Enums"
	"time"
)

func main() {

	education1 := statics.Education{
		Degree:      enums.Bachelors,
		Name:        "smthing",
		GradeSystem: enums.Percentage,
		Grade:       84.2,
	}
	education2 := statics.Education{
		Degree:      enums.Masters,
		Name:        "smthing",
		GradeSystem: enums.Percentage,
		Grade:       84.2,
	}

	eduArr := []statics.Education{education1, education2}
	tagArr := []enums.Tags{enums.Game, enums.Music, enums.Prog}

	myStruct := statics.Profile{
		Name:         "nucleo",
		Email:        "nucleo@gmail.com",
		BirthDate:    time.Now(),
		Education:    eduArr,
		LinkedIn:     statics.NewLink("linkedin", "https://linkedin.com/"),
		OtherLinks:   []statics.Link{statics.NewLink("github", "github.com/"), statics.NewLink("gmail", "smthing@mail.com")},
		Description:  "lorem ipsum",
		Followers:    []statics.UserReference{statics.UserReference{Id: 1}, statics.UserReference{Id: 2}},
		Following:    []statics.UserReference{statics.UserReference{Id: 1}, statics.UserReference{Id: 2}},
		TagsFollowed: tagArr,
	}

	data, _ := json.Marshal(myStruct)

	fmt.Println(string(data))
}
