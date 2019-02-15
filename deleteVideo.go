package main

import (
	"net/http"
)

func deleteVideo(writer http.ResponseWriter, request *http.Request) {
	// request.ParseForm()

	// err := db.Ping()
	// checkError(err)

	// if err != nil {

	// 	http.Redirect(writer, request, "/index", http.StatusSeeOther)

	// } else {

	// 	// Remove data from the table.
	// 	sqlStatement := "DELETE FROM videos WHERE url = $1;"

	// 	_, err = db.Exec(sqlStatement, "https://www.youtube.com/watch?v="+(request.PostForm["hidVideoID"][0]))
	// 	checkError(err)

	// 	http.Redirect(writer, request, "/index", http.StatusSeeOther)
	// }
}
