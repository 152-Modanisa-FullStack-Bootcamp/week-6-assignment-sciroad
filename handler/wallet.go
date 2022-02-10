package handler

import (
	"bootcamp/service"
	"encoding/json"
	"net/http"
)

type IWalletHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type WalletHandler struct {
	walletService service.IWalletService
}

func (h *WalletHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request"))
		return
	}
	wallets := h.walletService.GetAll()
	json, err := json.Marshal(wallets)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(json)

}
func (h *WalletHandler) Get(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[1:]
	wallet, err := h.walletService.Get(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json, err := json.Marshal(wallet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(json)

}
func (h *WalletHandler) Create(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[1:]
	if len(username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid username"))
		return
	}
	err := h.walletService.Create(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))
}
func (h *WalletHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[1:]
	if len(username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid username"))
		return
	}
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	amount, ok := body["balance"].(float64)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("balance is not a number"))
		return
	}
	err = h.walletService.Update(username, amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("updated"))
}
func NewWalletHandler(walletService service.IWalletService) IWalletHandler {
	return &WalletHandler{walletService}
}
