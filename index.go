package main

import (
	"database/sql"
	"net/http"
	"text/template"
)

func index(writer http.ResponseWriter, request *http.Request) {
	template, _ := template.ParseFiles("templates/index.html")

	err := db.Ping()
	checkError(err)

	if err != nil {
		template.Execute(writer, "Cannot connect to the database")
	} else {
		// Read data from table.
		var id int
		var name string
		var url string

		sqlStatement := "SELECT * FROM videos;"

		rows, err := db.Query(sqlStatement)
		checkError(err)

		defer rows.Close()

		videos := make(map[string]string)

		for rows.Next() {
			switch err := rows.Scan(&id, &name, &url); err {
			case sql.ErrNoRows:

				template.Execute(writer, "No data were returned")

			case nil:

				videos[url[32:len(url)]] = name

			default:

				checkError(err)

			}
		}

		template.Execute(writer, videos)
	}
}
