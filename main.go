package main

import (
	"net/http"
	"golang-main/bookstores.mb/middleware"
	"golang-main/bookstores.mb/routes"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.OrderRoutes(router)

	router.Use(middleware.TrackNumberOfRequests)

	http.ListenAndServe(":4040", router)
}