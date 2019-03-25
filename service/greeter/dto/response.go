package dto

type HelloResponse struct {
	SettingMessage string `json:"setting_message"`
	Id             string `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}
