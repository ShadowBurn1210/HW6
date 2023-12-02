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

func deleteBooksFromHTML(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookTitle := c.PostForm("title")

		err := deleteBook(db, bookTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func chooseBooksFromHTML(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.PostForm("pagesRead")
		pagesRead, _ := strconv.Atoi(a)
		progress := c.PostForm("progress")
		title := c.PostForm("title")
		err := updateBook(db, title, pagesRead, progress)
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

	router.GET("/delete", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deleteBook.html", gin.H{
			"title": "You can delete a book here!",
		})
	})
	router.POST("/delete", deleteBooksFromHTML(db))

	router.GET("/update", func(c *gin.Context) {
		c.HTML(http.StatusOK, "updateBook.html", gin.H{
			"title": "You can update a book here!",
		})
	})

	router.POST("/update", chooseBooksFromHTML(db))

	router.Run("localhost:8080")
	db.Close()

}
