package routes

import (
	"net/http"

	"jeanluc.freehzaix.com/resume/controllers"
)

func Web(){
	http.HandleFunc("/", controllers.Home)
}