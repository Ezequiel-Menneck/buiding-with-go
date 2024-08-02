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

type NoteToUser struct {
	NoteName     string
	Description  string
	CategoryName string
}

type Category struct {
	Id   int64
	Name string
}

func InsertNote(note Note) (*Note, error) {
	noteToReturn := &Note{}
	conn, err := OpenConection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	insertQuery := "INSERT INTO notes(note_name, description, category_id) VALUES ($1, $2, $3) RETURNING id, note_name, description, category_id;"

	err = conn.QueryRow(insertQuery, note.NoteName, note.Description, note.CategoryId).Scan(&noteToReturn.Id, &noteToReturn.NoteName, &noteToReturn.Description, &noteToReturn.CategoryId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return noteToReturn, nil

}

func GetNoteByName(noteName string) (*Note, error) {
	note := &Note{}

	conn, err := OpenConection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	selectQuery := `SELECT note_name, description, category_id FROM notes WHERE note_name = $1;`

	err = conn.QueryRow(selectQuery, noteName).Scan(&note.NoteName, &note.Description, &note.CategoryId)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func FindCategoryNameById(id int64) (string, error) {
	var categoryName string
	conn, err := OpenConection()
	if err != nil {
		return "", err

	}
	defer conn.Close()

	selectQuery := `SELECT name FROM categories WHERE id = $1;`

	err = conn.QueryRow(selectQuery, id).Scan(&categoryName)
	if err != nil {
		return "", err

	}

	return categoryName, nil

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
	var id int64

	conn, err := OpenConection()
	if err != nil {
		fmt.Println("error opening connection")
		return 0, err
	}
	defer conn.Close()

	insertQuery := "INSERT INTO categories (name) VALUES ($1) RETURNING id;"

	err = conn.QueryRow(insertQuery, categoryName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func FindAllCategories() ([]string, error) {
	var categories []string
	conn, err := OpenConection()
	if err != nil {
		fmt.Println("error opening connection")
		return nil, err
	}

	defer conn.Close()
	selectQuery := `SELECT name FROM categories;`

	rows, err := conn.Query(selectQuery)
	if err != nil {
		fmt.Println("error getting categories")
		return nil, err
	}

	for rows.Next() {
		var category string

		err = rows.Scan(&category)
		if err != nil {
			continue
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func FindAllNotesName() ([]string, error) {
	var notes []string
	conn, err := OpenConection()
	if err != nil {
		fmt.Println("error opening connection")
		return nil, err
	}

	defer conn.Close()
	selectQuery := `SELECT note_name FROM notes;`

	rows, err := conn.Query(selectQuery)
	if err != nil {
		fmt.Println("failed to get notes")
		return nil, err
	}

	for rows.Next() {
		var note string
		err = rows.Scan(&note)
		if err != nil {
			continue
		}

		notes = append(notes, note)
	}

	return notes, nil
}

type NoteUpdate struct {
	NoteName    string
	NewNoteName string
	Description string
	CategoryId  int64
}

func UpdateNote(note NoteUpdate) (int64, error) {
	conn, err := OpenConection()
	if err != nil {
		fmt.Println("error opening connection")
		return 0, err
	}

	defer conn.Close()

	sqlToUpdateQuery := `UPDATE notes SET (note_name, description, category_id) = ($1, $2, $3) WHERE note_name = $4;`

	result, err := conn.Exec(sqlToUpdateQuery, note.NewNoteName, note.Description, note.CategoryId, note.NoteName)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func FindAllNotes() ([]NoteToUser, error) {
	conn, err := OpenConection()
	if err != nil {
		fmt.Println("error opening connection")
		return nil, err
	}
	defer conn.Close()

	sqlGetAllNotes := `SELECT note_name, description, category_id FROM notes;`

	rows, err := conn.Query(sqlGetAllNotes)
	if err != nil {
		fmt.Println("error getting notes")
		return nil, err
	}

	var notesToUser []NoteToUser
	for rows.Next() {
		var noteToUser NoteToUser
		var categoryId int64

		err = rows.Scan(&noteToUser.NoteName, &noteToUser.Description, &categoryId)
		if err != nil {
			continue
		}

		categoryName, err := FindCategoryNameById(categoryId)
		if err != nil {
			fmt.Println("error getting category")
			continue
		}

		noteToUser.CategoryName = categoryName
		notesToUser = append(notesToUser, noteToUser)
	}

	return notesToUser, nil
}
