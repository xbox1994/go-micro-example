package dto

type HelloResponse struct {
	Environment string `json:"environment"`
	Message     string `json:"message"`
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
