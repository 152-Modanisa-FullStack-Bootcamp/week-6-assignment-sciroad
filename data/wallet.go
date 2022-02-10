package data

import (
	"bootcamp/model"
	"errors"
)

type IWalletRepository interface {
	Get(username string) (*model.Wallet, error)
	GetAll() []model.Wallet
	Create(username string, initialBalanceAmount float64) error
	Update(username string, amount float64) error
}
type WalletRepository struct {
	wallets map[string]float64
}

func (w *WalletRepository) Get(username string) (*model.Wallet, error) {
	balance, ok := w.wallets[username]
	if !ok {
		return nil, errors.New("user does not exist")
	}
	wallet := model.Wallet{Username: username, Balance: balance}
	return &wallet, nil
}
func (w *WalletRepository) GetAll() []model.Wallet {
	wallets := []model.Wallet{}
	for username, balance := range w.wallets {
		wallets = append(wallets, model.Wallet{Username: username, Balance: balance})
	}
	return wallets
}
func (w *WalletRepository) Create(username string, initialBalanceAmount float64) error {
	if _, ok := w.wallets[username]; ok {
		return errors.New("user already exists")
	}
	w.wallets[username] = initialBalanceAmount
	return nil
}
func (w *WalletRepository) Update(username string, amount float64) error {
	balance, ok := w.wallets[username]
	if !ok {
		return errors.New("user does not exist")
	}
	w.wallets[username] = balance + amount
	return nil
}
func NewWalletRepository() IWalletRepository {
	wallets := map[string]float64{}
	return &WalletRepository{
		wallets: wallets,
	}
}
func newWalletRepositoryForTest(wallets map[string]float64) IWalletRepository {
	return &WalletRepository{
		wallets: wallets,
	}
}
