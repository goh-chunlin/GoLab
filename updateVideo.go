package main

import (
	"net/http"
)

func updateVideo(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	err := db.Ping()
	checkError(err)

	if err != nil {

		http.Redirect(writer, request, "/index", http.StatusSeeOther)

	} else {

		// Update data in the table.
		sqlStatement := "UPDATE videos SET name = $1 WHERE url = $2;"

		_, err = db.Exec(sqlStatement, request.PostForm["VideoName"][0], "https://www.youtube.com/watch?v="+(request.PostForm["hidVideoID"][0]))
		checkError(err)

		http.Redirect(writer, request, "/index", http.StatusSeeOther)
	}
}
