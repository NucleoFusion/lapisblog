package profileRoute

import (
	"fmt"
	"net/http"
)

type addProfileStruct struct{}

func (_ *addProfileStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/profile/add route.")
}
