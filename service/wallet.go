package service

import (
	"bootcamp/data"
	"bootcamp/model"
	"errors"
)

type IWalletService interface {
	GetAll() []model.Wallet
	Get(username string) (*model.Wallet, error)
	Create(username string) error
	Update(username string, amount float64) error
}
type WalletService struct {
	repository           data.IWalletRepository
	initialBalanceAmount float64
	minimumBalanceAmount float64
}

func (w *WalletService) Get(username string) (*model.Wallet, error) {
	return w.repository.Get(username)
}
func (w *WalletService) Create(username string) error {
	return w.repository.Create(username, w.initialBalanceAmount)
}
func (w *WalletService) Update(username string, amount float64) error {
	wallet, err := w.repository.Get(username)
	if err != nil {
		return err
	}
	if wallet.Balance+amount < w.minimumBalanceAmount {
		return errors.New("invalid amount")
	}
	newBalane := wallet.Balance + amount
	err = w.repository.Update(username, newBalane)
	return err
}
func (w *WalletService) GetAll() []model.Wallet {
	return w.repository.GetAll()
}

func NewWalletService(walletRepository data.IWalletRepository, initialBalanceAmount, minimumBalanceAmount float64) IWalletService {
	return &WalletService{
		repository:           walletRepository,
		initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount,
	}
}
