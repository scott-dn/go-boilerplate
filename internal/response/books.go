package response

import "github.com/scott-dn/go-boilerplate/internal/database/entities"

type Books struct {
	Data []*entities.Book `json:"data"`
}

type Book struct {
	Data *entities.Book `json:"data"`
}
