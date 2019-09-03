package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rwirdemann/treubuilder/web"
)

func main() {
	web.AddAccountSystemHandler()
	web.AddAccountHandler()
	web.AddAccountHandlerVue()
	fs := http.FileServer(http.Dir("js"))
	web.R.PathPrefix("/").Handler(fs)

	web.R.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		fmt.Println(tpl)
		return nil
	})

	http.ListenAndServe(":8080", web.R)
}
