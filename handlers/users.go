package handlers

import (
	"fmt"
	"net/http"

	"github.com/switching/go-web-poc/dbService"
	"github.com/switching/go-web-poc/models"
)

//UserHandlers asdfas dasf
type UserHandlers struct {
	db dbService.UserDBService
}

//NewUserHandler asdasdasdsa
func NewUserHandler(db dbService.UserDBService) *UserHandlers {
	return &UserHandlers{db}
}

//Create adasdasdsad
func (u *UserHandlers) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating exmaple user")
	user1 := &models.User{
		Email:     "my@email.com",
		Firstname: "Jon",
		Middle:    "",
		Surname:   "Doe",
		Dob:       "01/01/1970",
	}
	user1.SetPassword("aaaaaaaaaa")
	fmt.Println(u.db.Create(user1))
}
