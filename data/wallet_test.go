package data

import (
	"bootcamp/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Get(t *testing.T) {
	testUser := "testUser1"
	testBalance := float64(100)
	wallets := map[string]float64{
		testUser: testBalance,
	}

	repository := newWalletRepositoryForTest(wallets)
	wallet, err := repository.Get(testUser)
	assert.Equal(t, wallet.Username, testUser)
	assert.Equal(t, wallet.Balance, testBalance)
	assert.Nil(t, err)
}
func TestRepository_GetUserDoesNotExists(t *testing.T) {
	testUser := "testUser1"
	testBalance := float64(100)
	wallets := map[string]float64{
		testUser: testBalance,
	}
	repository := newWalletRepositoryForTest(wallets)
	wallet, err := repository.Get("test")
	assert.Equal(t, err.Error(), "user does not exist")
	assert.Nil(t, wallet)
}
func TestRepository_GetAll(t *testing.T) {
	testUser := "testUser1"
	testBalance := float64(100)
	testUser2 := "testUser2"
	testBalance2 := float64(0)
	wallets := map[string]float64{
		testUser:  testBalance,
		testUser2: testBalance2,
	}
	expectedWallets := []model.Wallet{}
	for username, balance := range wallets {
		expectedWallets = append(expectedWallets, model.Wallet{Username: username, Balance: balance})
	}

	repository := newWalletRepositoryForTest(wallets)

	returnedWallets := repository.GetAll()
	for _, wallet := range returnedWallets {
		balance, ok := wallets[wallet.Username]
		assert.True(t, ok)
		assert.Equal(t, wallet.Balance, balance)
	}

}
func TestRepository_Create(t *testing.T) {
	testUser := "testUser1"
	testBalance := float64(100)
	repository := NewWalletRepository()
	err := repository.Create(testUser, testBalance)
	assert.Nil(t, err)
}
func TestRepository_CreateUserAlreadyExists(t *testing.T) {
	testUser := "testUser"
	testBalance := float64(100)
	wallets := map[string]float64{
		testUser: testBalance,
	}
	repository := newWalletRepositoryForTest(wallets)
	err := repository.Create(testUser, testBalance)
	assert.Equal(t, err.Error(), "user already exists")
}
func TestRepository_Update(t *testing.T) {
	testUser := "testUser"
	testBalance := float64(100)
	wallets := map[string]float64{
		testUser: testBalance,
	}
	repository := newWalletRepositoryForTest(wallets)
	err := repository.Update(testUser, testBalance)
	assert.Nil(t, err)
}
func TestRepository_UpdateUserDoesNotExists(t *testing.T) {
	testUser := "testUser"
	testBalance := float64(100)
	wallets := map[string]float64{}
	repository := newWalletRepositoryForTest(wallets)
	err := repository.Update(testUser, testBalance)
	assert.Equal(t, err.Error(), "user does not exist")
}
