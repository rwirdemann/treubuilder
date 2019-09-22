package main

import (
	"fmt"
	"net/http"

	"github.com/rwirdemann/treubuilder/web"
)

func main() {
	web.AddAccountSystemHandler()
	web.AddAccountHandler()
	web.AddAccountHandlerVue()
	fs := http.FileServer(http.Dir("js"))
	web.R.PathPrefix("/").Handler(fs)
	fmt.Printf("http://localhost:8080/accountsystem")
	http.ListenAndServe(":8080", web.R)
}
