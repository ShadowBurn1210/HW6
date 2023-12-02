package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	_ "strconv"
)

// https://www.youtube.com/watch?v=bj77B59nkTQ - video I used to help me with this assignment, mainly API stuff

type book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     int    `json:"pages"`
	PagesRead int    `json:"pagesRead"`
	Progress  string `json:"progress"`
}

func renderBooks(c *gin.Context, books []book) {
	c.HTML(http.StatusOK, "books.html", gin.H{"Books": books})
}

func showBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := retrieveBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		renderBooks(c, books)
	}
}

func addBooksFromHTML(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.PostForm("pages")
		pages, _ := strconv.Atoi(a)
		newBook := book{
			Title:     c.PostForm("title"),
			Author:    c.PostForm("author"),
			Pages:     pages,
			PagesRead: 0,
			Progress:  "Not Started",
		}
		err := addBooks(newBook, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func main() {
	db := ConnectToDatabase()

	router := gin.Default()
	router.LoadHTMLGlob("backend/templates/*")
	CreateBookTable(db)

	router.GET("/", func(c *gin.Context) {
		showBooks(db)(c)
	})
	router.GET("/addbook", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addbook.html", gin.H{
			"title": "You can add a new book here!",
		})
	})
	router.POST("/addbook", addBooksFromHTML(db))

	router.Run("localhost:8080")
	db.Close()

}

//func main() {
//	db := ConnectToDatabase()
//	CreateBookTable(db)
//
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Println("Enter some lines (Ctrl+D to end):")
//
//	for scanner.Scan() {
//		line := scanner.Text()
//		fmt.Println("Read:", line)
//
//		if line == "add" {
//			wheelOfTimeBook := book{
//				Title:     "The Eye of the World",
//				Author:    "Robert Jordan",
//				Pages:     814,
//				PagesRead: 200,
//				Progress:  "In Progress",
//			}
//			err := addBooks(wheelOfTimeBook, db)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		} else if line == "show" {
//			books, err := retrieveBooks(db)
//			if err != nil {
//				return
//			}
//			for _, book := range books {
//				fmt.Println(book)
//			}
//		} else if line == "delete" {
//			fmt.Print("Enter the ID of the book to delete: ")
//			var bookID int
//			_, err := fmt.Scan(&bookID)
//			if err != nil {
//				log.Println("Error reading book ID:", err)
//				continue
//			}
//
//			err = deleteBook(db, bookID)
//			if err != nil {
//				log.Println("Error deleting book:", err)
//			}
//		} else if line == "update" {
//			fmt.Print("Enter the ID of the book to update: ")
//			var bookID int
//			_, err := fmt.Scan(&bookID)
//			if err != nil {
//				log.Println("Error reading book ID:", err)
//				continue
//			}
//
//			wheelOfTimeBook := book{
//				Title:     "The Eye of the World",
//				Author:    "Robert Jordan",
//				Pages:     900,
//				PagesRead: 210,
//				Progress:  "In Progress",
//			}
//
//			err = updateBook(db, bookID, wheelOfTimeBook)
//			if err != nil {
//				log.Println("Error updating book:", err)
//			}
//
//		} else if line == "exit" {
//			break
//		}
//		db.Close()
//	}
//}
