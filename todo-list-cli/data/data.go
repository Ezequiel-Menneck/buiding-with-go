package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func OpenConection() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456789 dbname=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}

func CreateTableCategories() {
	createTableCategoriesSQL := `CREATE TABLE IF NOT EXISTS categories (
    	"id" SERIAL PRIMARY KEY,
    	"name" VARCHAR(255) NOT NULL
	);`

	conn, err := OpenConection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	statement, err := conn.Prepare(createTableCategoriesSQL)
	if err != nil {
		fmt.Errorf("someting went wrong trying to creating categories table")
	}

	exec, err := statement.Exec()
	if err != nil {
		return
	}
	fmt.Println("table created", exec)
}

func CreateTableNotes() {
	createTableNotesSQL := `CREATE TABLE IF NOT EXISTS notes (
		"id" SERIAL PRIMARY KEY,
		"note_name" VARCHAR(255) NOT NULL,
		"description" VARCHAR(255) NOT NULL,
		"category_id" integer NOT NULL,
        FOREIGN KEY ("category_id") REFERENCES categories("id")
	  );`

	conn, err := OpenConection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	statement, err := conn.Prepare(createTableNotesSQL)
	if err != nil {
		fmt.Errorf("someting went wrong trying to creating notes table")
	}

	exec, err := statement.Exec()
	if err != nil {
		return
	}
	fmt.Println("table created", exec)
}

type Note struct {
	Id          int64
	NoteName    string
	Description string
	CategoryId  int64
}

type Category struct {
	Id   int64
	Name string
}

func InsertNote(note Note) (noteToReturn Note, err error) {
	conn, err := OpenConection()
	if err != nil {
		return noteToReturn, err
	}
	defer conn.Close()

	insertQuery := "INSERT INTO notes(note_name, description, category_id) VALUES ($1, $2, $3) RETURNING id, note_name, description, category_id;"

	err = conn.QueryRow(insertQuery, note.NoteName, note.Description, note.CategoryId).Scan(&noteToReturn.Id, &noteToReturn.NoteName, &noteToReturn.Description, &noteToReturn.CategoryId)
	if err != nil {
		fmt.Println(err)
		return noteToReturn, err
	}

	return noteToReturn, nil

}

func GetNoteByName(noteName string) (string, error) {
	var noteNameDb string

	conn, err := OpenConection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	selectQuery := `SELECT note_name FROM notes WHERE note_name = $1;`

	err = conn.QueryRow(selectQuery, noteName).Scan(&noteNameDb)
	if err != nil {
		return "", err
	}

	return noteNameDb, nil
}

func FindCategoryByName(categoryName string) (int64, error) {
	var id int64
	conn, err := OpenConection()
	if err != nil {
		return 0, err

	}
	defer conn.Close()

	selectQuery := `SELECT id FROM categories WHERE name = $1;`

	err = conn.QueryRow(selectQuery, categoryName).Scan(&id)
	if err != nil {
		return 0, err

	}

	return id, nil

}

func CreateCategory(categoryName string) (int64, error) {
	return 0, nil
}
