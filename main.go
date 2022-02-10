package main

import (
	"bootcamp/config"
	"bootcamp/data"
	"bootcamp/handler"
	"bootcamp/service"
	"net/http"
)

func main() {
	config := config.Get()
	repository := data.NewWalletRepository()
	service := service.NewWalletService(repository, config.InitialBalanceAmount, config.MinumumBalanceAmount)
	handler := handler.NewWalletHandler(service)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if r.URL.Path == "/" {
			handler.GetAll(w, r)
			return
		}

		switch method {
		case "GET":
			handler.Get(w, r)
		case "POST":
			handler.Update(w, r)
		case "PUT":
			handler.Create(w, r)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("he çoğsundur"))
		}

	})
	http.ListenAndServe(":8080", nil)
}
