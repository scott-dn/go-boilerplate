package service

import (
	"github.com/scott-dn/go-boilerplate/internal/service/book"
	"gorm.io/gorm"
)

type Service struct {
	Book book.IBook
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		Book: book.NewBookService(db),
	}
}
