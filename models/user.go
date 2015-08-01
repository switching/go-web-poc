// Package models provides all the data models used within the system
package models

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//User holds current user of the system
type User struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email        string        `json:"email" bson:"email"`
	Firstname    string        `json:"firstname" bson:"firstname"`
	Middle       string        `json:"middle,omitempty" bson:"middle,omitempty"`
	Surname      string        `json:"surname" bson:"surname"`
	Dob          string        `json:"dob" bson:"dob"`
	passwordHash string
}

func (u *User) hasMiddle() bool {
	if u.Middle == "" {
		return true
	}
	return false
}

// Fullname combines, Firstname, Middle (if set) and Surname into one string
func (u *User) Fullname(capsFirstLetter bool) string {
	var fn string
	if u.hasMiddle() {
		fn = fmt.Sprintf("%s %s %s", u.Firstname, u.Middle, u.Surname)
	} else {
		fn = fmt.Sprintf("%s %s", u.Firstname, u.Surname)
	}

	if capsFirstLetter {
		fn = strings.Title(fn)
	}
	return fn
}

// InitialsWithSurname will return a string containng the first letter of
// the Firstname, Middle (if set) and the whole Surname
func (u *User) InitialsWithSurname() string {
	var n string
	if u.hasMiddle() {
		n = fmt.Sprintf("%s.%s %s", u.Firstname[0:1], u.Middle[0:1], u.Surname)
	} else {
		n = fmt.Sprintf("%s %s", u.Firstname[0:1], u.Surname)
	}
	n = strings.Title(n)
	return n
}

//SetPassword accepts the users password and hashes it, sets passwordHash
func (u *User) SetPassword(password string) error {
	log.Println("setting hashed password")
	u.passwordHash = "password hash"
	return nil
}

//PasswordHash gets the password hash for ths curent user
func (u *User) PasswordHash() string {
	return u.passwordHash
}
