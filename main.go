package main

import (
	"golang-main/bookstores.mb/initializers"
	"golang-main/bookstores.mb/middleware"
	"golang-main/bookstores.mb/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	router := mux.NewRouter()

	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.OrderRoutes(router)

	router.Use(middleware.TrackNumberOfRequests)

	http.ListenAndServe(":4040", router)
}