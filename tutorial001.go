package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main2() {
	// for _, e := range os.Environ() {
	// 	pair := strings.Split(e, "=")
	// 	fmt.Println(pair[0])
	// }

	// Initialize connection string.
	var connectionString = fmt.Sprintf(os.Getenv("CONNECTION_STRING"))

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	// Drop previous table of same name if one exists
	_, err = db.Exec("DROP TABLE IF EXISTS videos;")
	checkError(err)
	fmt.Println("Finished dropping table if existed")

	// Create table
	_, err = db.Exec("CREATE TABLE videos (id serial PRIMARY KEY, name VARCHAR(255), url VARCHAR(52));")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into the table.
	sqlStatement := "INSERT INTO videos (name, url) VALUES ($1, $2);"

	_, err = db.Exec(sqlStatement, "Fate/Stay Night: Unlimited Blade Works - Ideal White", "https://www.youtube.com/watch?v=LfItvEgyW9k")
	checkError(err)

	_, err = db.Exec(sqlStatement, "Lanota - Dreams Go On", "https://www.youtube.com/watch?v=X0qmks7lEyc")
	checkError(err)

	_, err = db.Exec(sqlStatement, "Steins:Gate - Last Game", "https://www.youtube.com/watch?v=SQBHr1kGmT0")
	checkError(err)

	_, err = db.Exec(sqlStatement, "Uchiage Hanabi - Fireworks", "https://www.youtube.com/watch?v=jcOKQRV0JJE")
	checkError(err)

	fmt.Println("Inserted four records")

	// Read data from table.
	var id int
	var name string
	var url string

	sqlStatement = "SELECT * FROM videos;"

	rows, err := db.Query(sqlStatement)
	checkError(err)

	defer rows.Close()

	fmt.Println("---------------------------------------------------------")
	fmt.Println("                     MyYouTube Data                      ")
	fmt.Println("---------------------------------------------------------")

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &url); err {
		case sql.ErrNoRows:
			fmt.Println("No data were returned")
		case nil:
			fmt.Printf(" ID: %d \n", id)
			fmt.Printf(" Video : %s \n", name)
			fmt.Printf(" URL: %s \n", url)
			fmt.Println("---------------------------------------------------------")
		default:
			checkError(err)
		}
	}

	var selectedID int

	fmt.Print("ID of the record that you want to play [0]: ")
	fmt.Scan(&selectedID)

	sqlStatement = "SELECT * FROM videos WHERE id=$1;"

	rows, err = db.Query(sqlStatement, selectedID)
	checkError(err)

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &url); err {
		case sql.ErrNoRows:
			fmt.Println("No data were returned")
		case nil:
			openbrowser(url)
			break
		default:
			checkError(err)
		}
	}

	defer rows.Close()
}
