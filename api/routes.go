package api

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/scott-dn/go-boilerplate/internal/app"
	"github.com/scott-dn/go-boilerplate/internal/request"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func registerRoutes(group *echo.Group, app *app.App) {
	// enforce authentication for all routes in this group
	group.Use(auth(app))

	group.GET("/books", getBooks(app))
	group.GET("/books/:bookId", getBookByID(app))
	group.POST("/books", addBook(app), requireAdmin)
	group.PUT("/books/:bookId", updateBook(app), requireAdmin)
	group.DELETE("/books/:bookId", deleteBook(app), requireAdmin)
}

// getBooks godoc
//
//	@Summary		Get all books
//	@Description	Get all books
//	@Tags			books
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.Books
//	@Router			/books [get]
func getBooks(app *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, app.Service.Book.GetBooks())
	}
}

// getBookByID godoc
//
//	@Summary		Get book by id
//	@Description	Get book by id
//	@Tags			books
//	@Security		ApiKeyAuth
//	@Param			bookId	path		string	true	"book id"
//	@Success		200		{object}	response.Book
//	@Router			/books/{bookId} [get]
func getBookByID(app *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		bookID := c.Param("bookId")
		id, err := strconv.ParseUint(bookID, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
		}
		res, err := app.Service.Book.GetBookByID(uint(id))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}

// addBook godoc
//
//	@Summary		Create book
//	@Description	Create book
//	@Tags			books
//	@Security		ApiKeyAuth
//	@Param			request	body		request.AddBook	true	"request body"
//	@Success		201		{object}	response.Book
//	@Router			/books [post]
func addBook(app *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.AddBook
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(&req); err != nil {
			return err
		}
		c.Set("body", &req)
		return c.JSON(http.StatusCreated, app.Service.Book.AddBook(&req, c.Get("email").(string))) //nolint:forcetypeassert
	}
}

// updateBook godoc
//
//	@Summary		Update book
//	@Description	Update book
//	@Tags			books
//	@Security		ApiKeyAuth
//	@Param			bookId	path		string				true	"book id"
//	@Param			request	body		request.UpdateBook	true	"request body"
//	@Success		200		{object}	response.Book
//	@Router			/books/{bookId} [put]
func updateBook(app *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UpdateBook
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(&req); err != nil {
			return err
		}
		bookID := c.Param("bookId")
		id, err := strconv.ParseUint(bookID, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
		}
		c.Set("body", &req)
		res, err := app.Service.Book.UpdateBook(uint(id), &req, c.Get("email").(string))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}

// deleteBook godoc
//
//	@Summary		Delete book
//	@Description	Delete book
//	@Tags			books
//	@Security		ApiKeyAuth
//	@Param			bookId	path		string	true	"book id"
//	@Success		200		{object}	response.Book
//	@Router			/books/{bookId} [delete]
func deleteBook(app *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		bookID := c.Param("bookId")
		id, err := strconv.ParseUint(bookID, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
		}
		res, err := app.Service.Book.DeleteBook(uint(id))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}
