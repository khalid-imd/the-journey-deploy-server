package routes

import (
	"backend-journey/handlers"
	"backend-journey/pkg/middleware"
	"backend-journey/pkg/mysql"
	"backend-journey/repositories"

	"github.com/gorilla/mux"
)

func BookmarkRoutes(r *mux.Router) {
	bookmarkRepository := repositories.RepositoryBookmark(mysql.DB)
	h := handlers.HandlerBookmark(bookmarkRepository)

	r.HandleFunc("/bookmark", middleware.Auth(h.CreateBookmark)).Methods("POST")
	r.HandleFunc("/bookmark/{id}", h.GetBookmark).Methods("GET")
	r.HandleFunc("/bookmarks", h.FindBookmarks).Methods("GET")
	r.HandleFunc("/bookmark/{id}", middleware.Auth(h.DeleteBookmark)).Methods("DELETE")
}
