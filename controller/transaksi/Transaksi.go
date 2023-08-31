package transaksi

import (
	"encoding/json"
	"link-aja/connection"

	"link-aja/models"
	"link-aja/response"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
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

func Transfer(w http.ResponseWriter, r *http.Request) {
	account_number := chi.URLParam(r, "number_account")

	decoder := json.NewDecoder(r.Body)
	var datarequest response.Transfer

	if err := decoder.Decode(&datarequest); err != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	if datarequest.AccountNumber == account_number {
		ResponseApi(w, http.StatusInternalServerError, nil, "You Cannot Transfer To Your Own Account.")
		return
	}

	if datarequest.Balance <= 0 {
		ResponseApi(w, http.StatusInternalServerError, nil, "The Transferred Balance Must Not Be Less Than 1.")
		return
	}

	var dataAccount models.Account
	data := DB.Where("account_number = ?", account_number).Find(&models.Account{}).Scan(&dataAccount)
	if data.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, data.Error.Error())
		return
	}

	if dataAccount.AccountNumber == "" {
		ResponseApi(w, http.StatusInternalServerError, nil, "Sender Account Number Not Found")
		return
	}

	var dataAccountTransfer models.Account
	data = DB.Where("account_number = ?", datarequest.AccountNumber).Find(&models.Account{}).Scan(&dataAccountTransfer)
	if data.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, data.Error.Error())
		return
	}

	if dataAccountTransfer.AccountNumber == "" {
		ResponseApi(w, http.StatusInternalServerError, nil, "Recipient Account Number Not Found")
		return
	}

	if dataAccount.Balance < datarequest.Balance {
		ResponseApi(w, http.StatusInternalServerError, nil, "Insufficient Balance.")
		return
	}

	requestSender := models.Account{
		Balance: dataAccount.Balance - datarequest.Balance,
	}

	requestRecipient := models.Account{
		Balance: dataAccountTransfer.Balance + datarequest.Balance,
	}

	checkSender := DB.Model(&models.Account{}).Where("account_number = ?", account_number).Updates(&requestSender)
	if checkSender.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, checkSender.Error.Error())
		return
	}

	checkRecipient := DB.Model(&models.Account{}).Where("account_number = ?", datarequest.AccountNumber).Updates(&requestRecipient)
	if checkRecipient.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, checkRecipient.Error.Error())
		return
	}

	ResponseApi(w, http.StatusCreated, nil, "Success Transfer")
}

func AddBalance(w http.ResponseWriter, r *http.Request) {
	account_number := chi.URLParam(r, "number_account")

	decoder := json.NewDecoder(r.Body)
	var datarequest response.TopUp

	if err := decoder.Decode(&datarequest); err != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	var dataAccount models.Account
	data := DB.Where("account_number = ?", account_number).Find(&models.Account{}).Scan(&dataAccount)
	if data.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, data.Error.Error())
		return
	}

	if dataAccount.AccountNumber == "" {
		ResponseApi(w, http.StatusInternalServerError, nil, "Account Number Not Found")
		return
	}

	request := models.Account{
		Balance: dataAccount.Balance + datarequest.Balance,
	}

	check := DB.Model(&models.Account{}).Where("account_number = ?", account_number).Updates(&request)
	if check.Error != nil {
		ResponseApi(w, http.StatusInternalServerError, nil, check.Error.Error())
		return
	}

	ResponseApi(w, http.StatusCreated, nil, "Success Top Up Balance")
}
