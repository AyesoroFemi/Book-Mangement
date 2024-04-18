package routes

import (
	"golang-main/bookstores.mb/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/signup", controllers.Signup)
}