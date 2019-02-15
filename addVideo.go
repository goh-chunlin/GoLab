package main

import (
	"net/http"
)

func addVideo(writer http.ResponseWriter, request *http.Request) {
	// request.ParseForm()

	// err := db.Ping()
	// checkError(err)

	// if err != nil {

	// 	http.Redirect(writer, request, "/index", http.StatusSeeOther)

	// } else {

	// 	// Insert data into the table.
	// 	sqlStatement := "INSERT INTO videos (name, url) VALUES ($1, $2);"

	// 	_, err = db.Exec(sqlStatement, request.PostForm["hidVideoName"][0], "https://www.youtube.com/watch?v="+(request.PostForm["hidVideoID"][0]))
	// 	checkError(err)

	// 	http.Redirect(writer, request, "/index", http.StatusSeeOther)
	// }
}
