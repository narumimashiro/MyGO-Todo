package go_modules

import (
	"fmt"
	"net/http"
)

func HandleAdd (w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// フォーム送信されてきたデータの取得
		task := r.FormValue("task")

		// for debug
		fmt.Println("New task : ", task)

		fmt.Fprint(w, "success add task")
	} else {
		http.Error(w, "method add task failed", http.StatusMethodNotAllowed)
	}
}