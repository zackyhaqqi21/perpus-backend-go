package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"perpusbackendgo/model"
	"perpusbackendgo/service"
)

type BookController struct {
	Service *service.BookService
}

func NewBookController(service *service.BookService) *BookController {
	return &BookController{Service: service}
}

func (ctrl *BookController) GetBooks(c *gin.Context) {
	books, err := ctrl.Service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data buku"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (ctrl *BookController) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := ctrl.Service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data tidak valid", "detail": err.Error()})
		return
	}
	if err := ctrl.Service.CreateBook(&book); err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Gagal membuat buku (kode mungkin sudah dipakai)"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (ctrl *BookController) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data tidak valid"})
		return
	}

	book, err := ctrl.Service.UpdateBook(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.Service.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}
	c.Status(http.StatusNoContent)
}
