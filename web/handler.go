package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rwirdemann/treubuilder/service"

	"github.com/rwirdemann/treubuilder/domain"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()
}

func AddAccountSystemHandler() {
	template := template.Must(template.ParseFiles("web/templates/account-system.html"))
	as := service.GetAccountSystem(1)
	R.HandleFunc("/accountsystem", func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, struct{ AccountSystem domain.AccountSystem }{as})
	})
}

func AddAccountHandler() {
	template := template.Must(template.ParseFiles("web/templates/account.html"))
	R.HandleFunc("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		a := service.GetAccount(id)
		template.Execute(w, struct{ Account domain.Account }{a})
	})
}

func AddAccountHandlerVue() {
	template := template.Must(template.ParseFiles("web/vue/account.html"))
	R.HandleFunc("/accounts.js/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		a := service.GetAccount(id)
		template.Execute(w, struct{ Account domain.Account }{a})
	})
}
