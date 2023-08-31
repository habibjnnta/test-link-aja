package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListAccount struct {
	AccountNumber string `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}

type CreateAcount struct {
	CustomerName string `json:"customer_name"`
}

type TopUp struct {
	Balance int `json:"balance"`
}

type Transfer struct {
	AccountNumber string `json:"to_account_number"`
	Balance       int    `json:"balance"`
}
