package account

import (
	"encoding/json"
	"fmt"

	"link-aja/connection"
	"link-aja/models"
	"link-aja/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	reservice := response.Response{}
	reservice.Code = code
	reservice.Message = msg
	reservice.Data = data
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json") //return type data nya

	json.NewEncoder(w).Encode(reservice)
}

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	var listCustomer []models.Customer
	DB.Preload(clause.Associations).Find(&listCustomer)

	ResponseApi(w, http.StatusOK, listCustomer, "Success Get All Data Customer")
}

func GetSaldo(w http.ResponseWriter, r *http.Request) {
	account_number := chi.URLParam(r, "number_account")

	err := DB.First(&models.Account{}, "account_number = ?", account_number).Error
	if err != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	var dataRequest models.Account
	DB.Where("account_number = ?", account_number).Find(&models.Account{}).Scan(&dataRequest)

	if dataRequest.AccountNumber == "" {
		ResponseApi(w, http.StatusInternalServerError, nil, "Account Number Not Found")
		return
	}

	ResponseApi(w, http.StatusOK, dataRequest, "Success Check Saldo Account")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest response.CreateAcount

	if err := decoder.Decode(&datarequest); err != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	customerNumber := ""
	accountNumber := ""

	var resultCustomer *models.Customer
	DB.Last(&models.Customer{}).Scan(&resultCustomer)
	if resultCustomer != nil {
		str := resultCustomer.CustomerNumber
		i, _ := strconv.Atoi(str)
		i = i + 1
		str = fmt.Sprintf("%07d", i)
		customerNumber = str
	} else {
		customerNumber = "0000010"
	}

	var resultAccount *models.Account
	DB.Last(&models.Account{}).Scan(&resultAccount)
	if resultAccount != nil {
		str := resultAccount.AccountNumber
		i, _ := strconv.Atoi(str)
		i = i + 1
		str = fmt.Sprintf("%07d", i)
		accountNumber = str
	} else {
		accountNumber = "0000001"
	}

	customer := models.Customer{
		CustomerNumber: customerNumber,
		Name:           datarequest.CustomerName,
		Accounts: models.Account{
			AccountNumber:  accountNumber,
			Balance:        0,
			CustomerNumber: customerNumber,
		},
	}

	if customerCreate := DB.Create(&customer); customerCreate.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, customerCreate.Error.Error())
		return
	}

	ResponseApi(w, http.StatusCreated, nil, "Success Create Account")
}
