package main

import (
	"net/http"
	"jeanluc.freehzaix.com/resume/routes"
)


func main(){

	// Définir la route /static comme repertoire statique où
	// les fichiers html, css, images etc seront stockés.
	fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	routes.Web() // Initialiser les routes

	// Démarrer le serveur sur le port 8080
	http.ListenAndServe(":8080", nil)

}