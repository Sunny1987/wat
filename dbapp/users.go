package dbapp

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

var (
	GetAllUsers               = "https://pgserv.herokuapp.com/users"
	Client      *resty.Client = resty.New()
)

//UserList will provide the list of users from database pgserv
type UserList []struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//GetUsers method will get the users list
func (u *UserList) GetUsers(l *log.Logger) {
	l.Println("**** Fetching the users****")
	resp, err := Client.R().SetHeader("Accept", "application/json").Get(GetAllUsers)
	if err != nil {
		l.Println(err)

	}
	json.Unmarshal(resp.Body(), u)

}
