package main

import (
	"net/http"

	"github.com/rwirdemann/treubuilder/web"
)

func main() {
	web.AddAccountSystemHandler()
	web.AddAccountHandler()
	web.AddAccountHandlerVue()
	fs := http.FileServer(http.Dir("js"))
	web.R.PathPrefix("/").Handler(fs)
	http.ListenAndServe(":8080", web.R)
}
