package admin

import (
	"fmt"
	"net/http"
)

func Get_user_data(w http.ResponseWriter, r *http.Request) {
	pathParam := r.PathValue("user_id")

	fmt.Print(pathParam)
}
