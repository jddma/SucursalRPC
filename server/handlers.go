package server

/*
*	Agregar en este fichero los handlers con dos parametros (http.ResponseWriter, *http.Request)
*/

import (
	"html/template"
	"net/http"
	"path"

	"../serverUtilities/auth"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	//Renderizar archivo html


	templatePath:= path.Join("views/", "index.html")
	template, _ := template.ParseFiles(templatePath)
	template.Execute(w, nil)

}

func HandlePanel(w http.ResponseWriter, r *http.Request) {


	context := auth.GetSessionData(r,cookieHandler)
	//Renderizar archivo html
	templatePath:= path.Join("views/", "panel.html")
	template, _ := template.ParseFiles(templatePath)
	template.Execute(w, context)

}