package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// https://chat.openai.com/share/391f4459-a727-43eb-9a55-590bca0e8396 - link to my OpenAI chat for db

// I mainly used go documentation in this file to help me with this assignment
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Kasjauns2003"
	dbname   = "HW6"
)

func ConnectToDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	CreateBookTable(db)

	return db
}

func CreateBookTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Books (
			id         SERIAL PRIMARY KEY,
			title      VARCHAR(128) NOT NULL,
			author     VARCHAR(255) NOT NULL,
		    pages      INT NOT NULL,
		    pagesRead  INT NOT NULL,
			progress      VARCHAR(128) NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
}

func addBooks(newBook book, db *sql.DB) error {
	sqlStatement := `
		INSERT INTO Books (title, author, pages, pagesRead, progress)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var bookID string
	err := db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Pages, newBook.PagesRead, newBook.Progress).Scan(&bookID)
	if err != nil {
		return err
	}

	fmt.Printf("New book has been inserted with ID %s\n", bookID)
	return nil
}

func retrieveBooks(db *sql.DB) ([]book, error) {
	sqlStatement := `SELECT * FROM Books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []book
	for rows.Next() {
		var id int
		var title string
		var author string
		var pages int
		var pagesRead int
		var progress string

		err = rows.Scan(&id, &title, &author, &pages, &pagesRead, &progress)
		if err != nil {
			return nil, err
		}

		books = append(books, book{Title: title, Author: author, Pages: pages, PagesRead: pagesRead, Progress: progress})
	}

	return books, nil
}

func deleteBook(db *sql.DB, title string) error {
	sqlStatement := `DELETE FROM Books WHERE title = $1`
	_, err := db.Exec(sqlStatement, title)
	if err != nil {
		return err
	}

	fmt.Println("Book has been deleted")
	return nil
}

func updateBook(db *sql.DB, Booktitle string, pagesReads int, prograss string) error {
	sqlStatement := `
		UPDATE Books
		SET pagesRead = $2, progress = $3
		WHERE title = $1;`
	_, err := db.Exec(sqlStatement, Booktitle, pagesReads, prograss)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		return err
	}

	fmt.Println("Book has been updated")
	return nil
}
