package handler

import (
	"bootcamp/mock"
	"bootcamp/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWalletHandler_Create(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	testUser := "testUser"
	service.EXPECT().
		Create(testUser).
		Return(nil).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPut, "/"+testUser, nil)
	w := httptest.NewRecorder()
	handler.Create(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
	assert.Equal(t, w.Body.String(), "created")
}
func TestWalletHandler_CreateInvalidUsername(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPut, "/", nil)
	w := httptest.NewRecorder()
	handler.Create(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), "invalid username")
}

func TestWalletHandler_CreateUserAlreadyExists(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	testUser := "testUser"
	expectedError := errors.New("user does not exist")
	service.EXPECT().
		Create(testUser).
		Return(expectedError).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPut, "/"+testUser, nil)
	w := httptest.NewRecorder()
	handler.Create(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, w.Body.String(), expectedError.Error())
}
func TestWalletHandler_GetAll(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	serviceReturn := []model.Wallet{
		model.Wallet{
			Username: "testUser1",
			Balance:  float64(100),
		},
		model.Wallet{
			Username: "testUser2",
			Balance:  float64(0),
		},
	}
	service.EXPECT().
		GetAll().
		Return(serviceReturn).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAll(w, r)
	json, _ := json.Marshal(serviceReturn)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, w.Body.Bytes(), json)
}
func TestWalletHandler_GetAllMethodCheck(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPut, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAll(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), "Invalid Request")
}
func TestWalletHandler_GetAllEmptyList(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	serviceReturn := []model.Wallet{}
	service.EXPECT().
		GetAll().
		Return(serviceReturn).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAll(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, w.Body.String(), "[]")
}
func TestWalletHandler_Get(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	testUser := "testUser"
	serviceReturn := &model.Wallet{
		Username: testUser,
		Balance:  float64(100),
	}
	service.EXPECT().
		Get(testUser).
		Return(serviceReturn, nil).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/"+testUser, nil)
	w := httptest.NewRecorder()
	handler.Get(w, r)
	json, _ := json.Marshal(serviceReturn)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, w.Body.Bytes(), json)
}
func TestWalletHandler_GetUserDoesNotExists(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	testUser := "testUser"
	expectedError := errors.New("user does not exist")
	service.EXPECT().
		Get(testUser).
		Return(nil, expectedError).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/"+testUser, nil)
	w := httptest.NewRecorder()
	handler.Get(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, w.Body.String(), expectedError.Error())
}
func TestWalletHandler_Update(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	testUser := "testUser"
	testBalance := float64(-100)
	service.EXPECT().
		Update(testUser, testBalance).
		Return(nil).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(fmt.Sprintf("{\"balance\":%f}", testBalance))))
	w := httptest.NewRecorder()
	handler.Update(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, w.Body.String(), "updated")
}
func TestWalletHandler_UpdateInvalidInput(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	testUser := "testUser"

	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte("asdasdasd")))
	w := httptest.NewRecorder()
	handler.Update(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), "invalid character 'a' looking for beginning of value")
}

func TestWalletHandler_UpdateUserDoesNotExist(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))

	testUser := "testUser"
	testBalance := float64(-100)
	expectedError := errors.New("user does not exist")
	service.EXPECT().
		Update(testUser, testBalance).
		Return(expectedError).
		Times(1)
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(fmt.Sprintf("{\"balance\": %f}", testBalance))))
	w := httptest.NewRecorder()
	handler.Update(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, w.Body.String(), "user does not exist")
}
func TestWalletHandler_UpdateInvalidBalanceType(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))

	testUser := "testUser"

	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(fmt.Sprintf("{\"balance\": \"asdasdasd\"}"))))
	w := httptest.NewRecorder()
	handler.Update(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), "balance is not a number")
}
func TestWalletHandler_UpdateInvalidUsername(t *testing.T) {
	service := mock.NewMockIWalletService(gomock.NewController(t))
	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	handler.Update(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), "invalid username")
}
