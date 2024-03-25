package types

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Erorr   interface{} `json:"error"`
}
