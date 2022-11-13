package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	JourneyRoutes(r)
	AuthRoutes(r)
	BookmarkRoutes(r)
}
