package routes

import (
	"golang-main/bookstores.mb/controllers"
	"golang-main/bookstores.mb/middleware"
	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orderBook/{bookId}", middleware.ValidateUser(controllers.OrderBook))
	router.HandleFunc("/listOrderedBooks", middleware.ValidateUser(controllers.ListAllOrderedBooks))
}
