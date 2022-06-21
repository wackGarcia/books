﻿package books

import (
	"database/sql"
	"fmt"
	"net/http"

	bookRepo "github.com/wackGarcia/books/data/book"
	userRepo "github.com/wackGarcia/books/data/user"
	"github.com/wackGarcia/books/domain/book"
	"github.com/wackGarcia/books/domain/user"
	handler "github.com/wackGarcia/books/http"
)

func HttpServer(port int, db *sql.DB) error {
	services := injectServices(db)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), services.Http())
}

func injectServices(db *sql.DB) *handler.Handler {
	return &handler.Handler{
		User: user.NewUserService(userRepo.New(db)),
		Book: book.NewBookService(bookRepo.New(db)),
	}
}
