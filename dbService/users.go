package dbService

import (
	"fmt"

	"github.com/switching/go-web-poc/models"
	"gopkg.in/mgo.v2"
)

//UserDBService interface describes the requirements for the data layer
type UserDBService interface {
	Create(user *models.User) string
	Read(email string) models.User
	// Update(user models.User)
	// Delete(user models.User)
}

// UserDBServiceMongo provides a db warapper to access info about the user
type UserDBServiceMongo struct {
	dbSession  *mgo.Session
	dbName     string
	collection string
}

//New will return a configured UserDBServiceMongo struct
func New(mgoSession *mgo.Session, dbName string, collection string) *UserDBServiceMongo {
	return &UserDBServiceMongo{
		dbSession:  mgoSession,
		dbName:     dbName,
		collection: collection,
	}
}

//Create allows the creation of a user object in the database
func (u *UserDBServiceMongo) Create(user *models.User) string {

	//Do some stuff for testing

	// get copy of the mongo session from the pool, then close on function exit
	ses := u.dbSession.Copy()
	defer ses.Close()

	c := ses.DB(u.dbName).C(u.collection)

	co, err := c.Count()
	if err != nil {
		//return usful info
	}

	fmt.Println(co)
	err = c.Insert(user)

	return fmt.Sprint(user)
}

//Read asdasdsdsada
func (u *UserDBServiceMongo) Read(email string) models.User {
	return models.User{Email: "another user", Firstname: "mary", Middle: "J", Surname: "car", Dob: "10/01/1970"}
}
