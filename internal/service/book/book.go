package book

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scott-dn/go-boilerplate/internal/database/entities"
	"github.com/scott-dn/go-boilerplate/internal/database/query"
	"github.com/scott-dn/go-boilerplate/internal/request"
	"github.com/scott-dn/go-boilerplate/internal/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBook interface {
	GetBooks() *response.Books
	GetBookByID(uint) (*response.Book, error)
	AddBook(*request.AddBook, string) *response.Book
	UpdateBook(uint, *request.UpdateBook, string) (*response.Book, error)
	DeleteBook(uint) (*response.Book, error)
}

type Book struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *Book {
	return &Book{db: db}
}

func (b *Book) GetBooks() *response.Books {
	entities, err := query.Book.Find()
	if err != nil {
		panic(err)
	}
	return &response.Books{Data: entities}
}

func (b *Book) GetBookByID(id uint) (*response.Book, error) {
	q := query.Book
	entity, err := q.Where(q.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "book id not found")
		}
		panic(err)
	}
	return &response.Book{Data: entity}, nil
}

func (b *Book) AddBook(req *request.AddBook, by string) *response.Book {
	q := query.Book
	entity := &entities.Book{
		Name:        req.Name,
		Author:      req.Author,
		Description: req.Description,
		Version:     1,
		CreatedBy:   by,
		UpdatedBy:   by,
	}
	err := q.Create(entity)
	if err != nil {
		panic(err)
	}
	return &response.Book{Data: entity}
}

func (b *Book) UpdateBook(id uint, req *request.UpdateBook, by string) (*response.Book, error) {
	q := query.Book
	count, err := q.Where(q.ID.Eq(id), q.Version.Eq(req.Version)).Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "book id not found")
	}
	entity := &entities.Book{
		ID:          id,
		Name:        req.Name,
		Author:      req.Author,
		Description: req.Description,
		Version:     req.Version + 1,
		UpdatedBy:   by,
	}
	if _, err := q.Returning(entity).Updates(entity); err != nil {
		panic(err)
	}
	return &response.Book{Data: entity}, nil
}

func (b *Book) DeleteBook(id uint) (*response.Book, error) {
	q := query.Book
	count, err := q.Where(q.ID.Eq(id)).Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "book id not found")
	}
	entities := []*entities.Book{}
	if result := b.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&entities); result.Error != nil {
		panic(result.Error)
	}
	return &response.Book{Data: entities[0]}, nil
}
