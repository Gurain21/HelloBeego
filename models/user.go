package models

type User struct {
	Name    string `json:"name"`
	Birtday string `json:"birthday"`
	Address string `json:"address"`
	Nick    string `json:"nick"`
}
