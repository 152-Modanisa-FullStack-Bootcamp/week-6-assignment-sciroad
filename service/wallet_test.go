package service

import (
	"bootcamp/mock"
	"bootcamp/model"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Create(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	initialBalanceAmount := float64(0)
	mockRepository.EXPECT().Create(testUser, initialBalanceAmount).Return(nil)
	service := NewWalletService(mockRepository, initialBalanceAmount, 0)
	err := service.Create(testUser)
	assert.Nil(t, err)
}

func TestService_CreateUserAlreadyExist(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	initialBalanceAmount := float64(0)
	expectedError := errors.New("User Already Exist")
	mockRepository.EXPECT().Create(testUser, initialBalanceAmount).Return(expectedError)
	service := NewWalletService(mockRepository, initialBalanceAmount, 0)
	err := service.Create(testUser)
	assert.ErrorIs(t, err, expectedError)
}
func TestService_Update(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	balance := float64(-200)
	wallet := &model.Wallet{Username: testUser, Balance: 500}
	mockRepository.EXPECT().Get(testUser).Return(wallet, nil)
	mockRepository.EXPECT().Update(testUser, float64(300)).Return(nil)
	service := NewWalletService(mockRepository, 0, 0)
	err := service.Update(testUser, balance)
	assert.Nil(t, err)
}
func TestService_Update_UserDoesNotExist(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	balance := float64(0)
	expectedError := errors.New("user does not exist")
	mockRepository.EXPECT().Get(testUser).Return(nil, expectedError)
	service := NewWalletService(mockRepository, 0, 0)
	err := service.Update(testUser, balance)
	assert.ErrorIs(t, err, expectedError)
}
func TestService_Update_InvalidAmount(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	balance := float64(-200)
	wallet := &model.Wallet{Username: testUser, Balance: 0}
	mockRepository.EXPECT().Get(testUser).Return(wallet, nil)
	service := NewWalletService(mockRepository, 0, 0)
	err := service.Update(testUser, balance)
	assert.Equal(t, err.Error(), "invalid amount")
}
func TestService_GetUserDoesNotExist(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	expectedError := errors.New("user does not exist")
	mockRepository.EXPECT().Get(testUser).Return(nil, expectedError)
	service := NewWalletService(mockRepository, 0, 0)
	wallet, err := service.Get(testUser)
	assert.ErrorIs(t, err, expectedError)
	assert.Nil(t, wallet)
}
func TestService_Get(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	expectedWallet := &model.Wallet{Username: testUser, Balance: 0}
	mockRepository.EXPECT().Get(testUser).Return(expectedWallet, nil)
	service := NewWalletService(mockRepository, 0, 0)
	wallet, err := service.Get(testUser)
	assert.Nil(t, err, nil)
	assert.Equal(t, wallet, expectedWallet)
}
func TestService_GetAll(t *testing.T) {
	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	testUser := "testUser"
	expectedWallets := []model.Wallet{model.Wallet{Username: testUser, Balance: 0}}
	mockRepository.EXPECT().GetAll().Return(expectedWallets)
	service := NewWalletService(mockRepository, 0, 0)
	wallets := service.GetAll()
	assert.Equal(t, wallets, expectedWallets)
}
