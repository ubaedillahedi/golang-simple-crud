package main

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// DB Connection
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DB connection failed")
	}

	// Auto Migration
	// db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	// bookRepositoryFile := book.NewRepositoryFile()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title:", book.Title)
	// }

	// book, err := bookRepository.FindByID(2)

	// fmt.Println("Title:", book.Title)

	// book := book.Book{
	// 	Title: "S100 Startup",
	// 	Description: "Good Book",
	// 	Price: 950000,
	// 	Rating: 4,
	// 	Discount: 0,
	// }

	// bookRepository.Create(book)

	// CRUD
	// ============
	// CREATE data
	// ============
	// book := book.Book{}
	// book.Title = "Atomic Habit"
	// book.Price = 120000
	// book.Discount = 12
	// book.Rating = 4
	// book.Description = "Buku self development tentang membangun kebiasaan baik dan menghilangkan kebiasaan buruk"
	
	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// ============
	// READ data
	// ============
	// var books []book.Book
	// err = db.Debug().Find(&books).Error
	// err = db.Where("title LIKE ?", "%Habit%").Find(&books).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object", b)
		
	// }

	// ============
	// UPDATE data
	// ============
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }
	// book.Title = "Man Tiger (Revised Edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }

	// ============
	// DELETE data
	// ============
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error deleting book record")
	// 	fmt.Println("==========================")
	// }

	
	// Routing
	router := gin.Default()
	
	v1 := router.Group("/v1")
 
	v1.GET("/", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Welcome Web API GO",
		})
	})
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id/:title", bookHandler.BookHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.Update)
	v1.DELETE("/books/:id", bookHandler.Delete)
	router.Run(":8888")
}

